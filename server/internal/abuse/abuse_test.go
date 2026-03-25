package abuse

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"
)

type memoryRecorder struct {
	events []Event
	err    error
}

func (r *memoryRecorder) RecordAbuseEvent(_ context.Context, event Event) error {
	if r.err != nil {
		return r.err
	}

	r.events = append(r.events, event)
	return nil
}

type memoryLogger struct {
	lines []string
}

func (l *memoryLogger) Printf(format string, args ...any) {
	l.lines = append(l.lines, strings.TrimSpace(fmt.Sprintf(format, args...)))
}

func TestDetectObservationFlagsPathProbing(t *testing.T) {
	t.Parallel()

	findings := DetectObservation(Observation{
		RequestID: "req-1",
		Method:    "GET",
		URL:       "https://example.com/admin/.env/wp-login.php",
		Path:      "/admin/.env/wp-login.php",
	})

	if len(findings) != 1 {
		t.Fatalf("expected one finding, got %d", len(findings))
	}

	finding := findings[0]
	if finding.RuleKey != "path-probing" {
		t.Fatalf("expected rule key path-probing, got %q", finding.RuleKey)
	}

	if !finding.Block {
		t.Fatalf("expected path probing finding to be block-worthy")
	}

	if finding.Severity != SeverityHigh {
		t.Fatalf("expected high severity, got %s", finding.Severity)
	}
}

func TestDetectObservationFlagsSpamRelayPayload(t *testing.T) {
	t.Parallel()

	findings := DetectObservation(Observation{
		RequestID:   "req-2",
		Method:      "POST",
		URL:         "https://example.com/send",
		ContentType: "application/json",
		Headers: map[string]string{
			"subject": "bulk newsletter campaign",
		},
		Body: `{
			"recipients": [
				"a@example.com",
				"b@example.com",
				"c@example.com",
				"d@example.com",
				"e@example.com",
				"f@example.com",
				"g@example.com",
				"h@example.com",
				"i@example.com",
				"j@example.com"
			],
			"mode": "relay"
		}`,
	})

	if len(findings) != 1 {
		t.Fatalf("expected one finding, got %d", len(findings))
	}

	finding := findings[0]
	if finding.RuleKey != "spam-relay" {
		t.Fatalf("expected rule key spam-relay, got %q", finding.RuleKey)
	}

	if !finding.Block {
		t.Fatalf("expected spam relay finding to be block-worthy")
	}

	if finding.Severity != SeverityCritical {
		t.Fatalf("expected critical severity, got %s", finding.Severity)
	}
}

func TestDetectBatchFlagsFanOutAndBlockedRetries(t *testing.T) {
	t.Parallel()

	findings := DetectBatch(BatchObservation{
		Targets: []string{
			"https://a.example.com",
			"https://b.example.com",
			"https://c.example.com",
			"https://d.example.com",
			"https://e.example.com",
			"https://f.example.com",
			"https://g.example.com",
		},
		BlockedTargets: []string{
			"127.0.0.1",
			"127.0.0.1",
			"169.254.169.254",
			"169.254.169.254",
		},
	})

	if len(findings) != 2 {
		t.Fatalf("expected two findings, got %d", len(findings))
	}

	if findings[0].RuleKey != "target-fan-out" {
		t.Fatalf("expected first finding to be target fan-out, got %q", findings[0].RuleKey)
	}

	if findings[1].RuleKey != "blocked-target-retry" {
		t.Fatalf("expected second finding to be blocked-target-retry, got %q", findings[1].RuleKey)
	}
}

func TestReporterRecordsBlockedFindings(t *testing.T) {
	t.Parallel()

	recorder := &memoryRecorder{}
	logger := &memoryLogger{}
	reporter := Reporter{
		Logger:   logger,
		Recorder: recorder,
		Now: func() time.Time {
			return time.Date(2026, 3, 25, 12, 0, 0, 0, time.UTC)
		},
		NewID: func() string {
			return "event-1"
		},
	}

	events, err := reporter.Handle(context.Background(), Observation{
		RequestID: "req-3",
		UserID:    stringPtr("user-1"),
		SourceIP:  stringPtr("203.0.113.10"),
		Method:    "POST",
		URL:       "https://example.com/admin/.env",
		Host:      "example.com",
		Path:      "/admin/.env",
		Target:    "example.com",
	}, []Finding{
		{
			RuleKey:  "path-probing",
			Category: CategoryScan,
			Severity: SeverityHigh,
			Message:  "request matches common path probing patterns",
			Block:    true,
			Evidence: []string{"/admin", "/.env"},
		},
		{
			RuleKey:  "log-only",
			Category: CategorySuspicious,
			Severity: SeverityLow,
			Message:  "something worth logging",
			Block:    false,
		},
	})
	if err != nil {
		t.Fatalf("unexpected reporter error: %v", err)
	}

	if len(events) != 1 {
		t.Fatalf("expected one recorded event, got %d", len(events))
	}

	if len(recorder.events) != 1 {
		t.Fatalf("expected recorder to receive one event, got %d", len(recorder.events))
	}

	event := recorder.events[0]
	if event.ID != "event-1" {
		t.Fatalf("expected event ID event-1, got %q", event.ID)
	}

	if event.ActionTaken != ActionBlocked {
		t.Fatalf("expected blocked action, got %s", event.ActionTaken)
	}

	if event.RuleKey != "path-probing" {
		t.Fatalf("expected path-probing rule key, got %q", event.RuleKey)
	}

	if len(logger.lines) != 2 {
		t.Fatalf("expected two log lines, got %d", len(logger.lines))
	}

	if !strings.Contains(logger.lines[0], "path-probing") {
		t.Fatalf("expected log output to include rule key, got %q", logger.lines[0])
	}
}

func TestReporterPropagatesRecorderError(t *testing.T) {
	t.Parallel()

	reporter := Reporter{
		Recorder: &memoryRecorder{err: errors.New("write failed")},
		NewID: func() string {
			return "event-2"
		},
	}

	_, err := reporter.Handle(context.Background(), Observation{
		RequestID: "req-4",
		Target:    "example.com",
	}, []Finding{
		{
			RuleKey:  "path-probing",
			Category: CategoryScan,
			Severity: SeverityHigh,
			Message:  "request matches common path probing patterns",
			Block:    true,
		},
	})
	if err == nil {
		t.Fatalf("expected recorder error to propagate")
	}
}
