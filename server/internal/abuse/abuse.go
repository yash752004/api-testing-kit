package abuse

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Severity string

const (
	SeverityLow      Severity = "low"
	SeverityMedium   Severity = "medium"
	SeverityHigh     Severity = "high"
	SeverityCritical Severity = "critical"
)

type Category string

const (
	CategoryScan        Category = "scan"
	CategorySpamRelay   Category = "spam_relay"
	CategorySuspicious  Category = "suspicious_pattern"
	CategoryBlockedHost Category = "blocked_host"
)

type Action string

const (
	ActionLogged  Action = "logged"
	ActionBlocked Action = "blocked"
)

type Observation struct {
	RequestID      string
	UserID         *string
	SessionID      *string
	SourceIP       *string
	Method         string
	URL            string
	Host           string
	Path           string
	Query          string
	Headers        map[string]string
	Body           string
	ContentType    string
	BlockedTarget  string
	Target         string
	Redirects      []string
	WindowRequests []string
	WindowTargets  []string
}

type BatchObservation struct {
	UserID             *string
	SessionID          *string
	SourceIP           *string
	RequestIDs         []string
	Targets            []string
	BlockedTargets     []string
	FailureCount       int
	WindowDuration     time.Duration
	RepeatedCategories []string
}

type Finding struct {
	RuleKey  string
	Category Category
	Severity Severity
	Message  string
	Evidence []string
	Tags     []string
	Block    bool
}

type Event struct {
	ID          string          `json:"id"`
	UserID      *string         `json:"userId,omitempty"`
	SessionID   *string         `json:"sessionId,omitempty"`
	RequestID   *string         `json:"requestId,omitempty"`
	SourceIP    *string         `json:"sourceIp,omitempty"`
	Target      *string         `json:"target,omitempty"`
	RuleKey     string          `json:"ruleKey"`
	Category    Category        `json:"category"`
	Severity    Severity        `json:"severity"`
	ActionTaken Action          `json:"actionTaken"`
	Message     string          `json:"message"`
	Details     json.RawMessage `json:"details,omitempty"`
	CreatedAt   time.Time       `json:"createdAt"`
}

type Logger interface {
	Printf(format string, args ...any)
}

type Recorder interface {
	RecordAbuseEvent(ctx context.Context, event Event) error
}

type Reporter struct {
	Logger   Logger
	Recorder Recorder
	Now      func() time.Time
	NewID    func() string
}

var (
	scanTokens = []string{
		"/.env",
		"/.git",
		"/admin",
		"/cgi-bin",
		"/debug",
		"/metrics",
		"/phpmyadmin",
		"/server-status",
		"/wp-admin",
		"/wp-login.php",
	}
	spamTokens = []string{
		"bcc",
		"bulk",
		"campaign",
		"mail",
		"newsletter",
		"recipient",
		"relay",
		"smtp",
		"unsubscribe",
	}
	emailPattern = regexp.MustCompile(`(?i)\b[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}\b`)
)

func DetectObservation(obs Observation) []Finding {
	findings := make([]Finding, 0, 2)

	if finding, ok := detectPathProbing(obs); ok {
		findings = append(findings, finding)
	}

	if finding, ok := detectSpamRelay(obs); ok {
		findings = append(findings, finding)
	}

	return findings
}

func DetectBatch(obs BatchObservation) []Finding {
	findings := make([]Finding, 0, 2)

	if finding, ok := detectTargetFanOut(obs); ok {
		findings = append(findings, finding)
	}

	if finding, ok := detectRepeatedBlockedTargets(obs); ok {
		findings = append(findings, finding)
	}

	return findings
}

func (r Reporter) Handle(ctx context.Context, obs Observation, findings []Finding) ([]Event, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	now := r.now()
	events := make([]Event, 0, len(findings))
	for _, finding := range findings {
		if r.Logger != nil {
			r.Logger.Printf(
				"abuse %s severity=%s rule=%s target=%s request_id=%s",
				finding.Category,
				finding.Severity,
				finding.RuleKey,
				firstNonEmpty(obs.Target, obs.Host, obs.URL),
				obs.RequestID,
			)
		}

		if !finding.Block {
			continue
		}

		event := Event{
			ID:          r.newID(),
			UserID:      obs.UserID,
			SessionID:   obs.SessionID,
			RequestID:   stringPtr(obs.RequestID),
			SourceIP:    obs.SourceIP,
			Target:      stringPtr(firstNonEmpty(obs.BlockedTarget, obs.Target, obs.Host, obs.URL)),
			RuleKey:     finding.RuleKey,
			Category:    finding.Category,
			Severity:    finding.Severity,
			ActionTaken: ActionBlocked,
			Message:     finding.Message,
			CreatedAt:   now,
		}

		details, err := marshalDetails(obs, finding)
		if err != nil {
			return nil, err
		}
		event.Details = details

		if r.Recorder != nil {
			if err := r.Recorder.RecordAbuseEvent(ctx, event); err != nil {
				return nil, err
			}
		}

		events = append(events, event)
	}

	return events, nil
}

func (r Reporter) now() time.Time {
	if r.Now != nil {
		return r.Now()
	}

	return time.Now().UTC()
}

func (r Reporter) newID() string {
	if r.NewID != nil {
		return r.NewID()
	}

	var raw [8]byte
	if _, err := rand.Read(raw[:]); err != nil {
		return fmt.Sprintf("abuse-%d", time.Now().UTC().UnixNano())
	}

	return hex.EncodeToString(raw[:])
}

func detectPathProbing(obs Observation) (Finding, bool) {
	joined := strings.ToLower(strings.Join([]string{obs.URL, obs.Path, obs.Query, obs.Body}, "\n"))
	matched := uniqueMatches(joined, scanTokens)
	if len(matched) == 0 {
		return Finding{}, false
	}

	severity := SeverityMedium
	block := false
	if len(matched) >= 3 {
		severity = SeverityHigh
		block = true
	}

	return Finding{
		RuleKey:  "path-probing",
		Category: CategoryScan,
		Severity: severity,
		Message:  "request matches common path probing patterns",
		Evidence: matched,
		Tags:     []string{"scan", "path-probing"},
		Block:    block,
	}, true
}

func detectSpamRelay(obs Observation) (Finding, bool) {
	joined := strings.ToLower(strings.Join(collectText(obs), "\n"))
	emailCount := len(emailPattern.FindAllString(joined, -1))
	matched := uniqueMatches(joined, spamTokens)
	if emailCount == 0 && len(matched) == 0 {
		return Finding{}, false
	}

	if emailCount < 3 && len(matched) == 0 {
		return Finding{}, false
	}

	severity := SeverityMedium
	block := false
	if emailCount >= 10 || len(matched) >= 3 {
		severity = SeverityCritical
		block = true
	}

	evidence := append([]string{fmt.Sprintf("emails:%d", emailCount)}, matched...)
	return Finding{
		RuleKey:  "spam-relay",
		Category: CategorySpamRelay,
		Severity: severity,
		Message:  "request content resembles spam relay behavior",
		Evidence: evidence,
		Tags:     []string{"spam", "relay"},
		Block:    block,
	}, true
}

func detectTargetFanOut(obs BatchObservation) (Finding, bool) {
	targets := uniqueStrings(obs.Targets)
	if len(targets) < 6 {
		return Finding{}, false
	}

	severity := SeverityHigh
	block := true
	if len(targets) >= 12 {
		severity = SeverityCritical
	}

	return Finding{
		RuleKey:  "target-fan-out",
		Category: CategoryScan,
		Severity: severity,
		Message:  "batch contains broad target fan-out consistent with scanning",
		Evidence: targets,
		Tags:     []string{"scan", "fan-out"},
		Block:    block,
	}, true
}

func detectRepeatedBlockedTargets(obs BatchObservation) (Finding, bool) {
	if len(obs.BlockedTargets) < 3 {
		return Finding{}, false
	}

	targets := uniqueStrings(obs.BlockedTargets)
	sort.Strings(targets)

	return Finding{
		RuleKey:  "blocked-target-retry",
		Category: CategoryBlockedHost,
		Severity: SeverityHigh,
		Message:  "repeated attempts against blocked targets were observed",
		Evidence: targets,
		Tags:     []string{"blocked", "retry"},
		Block:    true,
	}, true
}

func marshalDetails(obs Observation, finding Finding) (json.RawMessage, error) {
	payload := map[string]any{
		"finding": finding,
		"observation": map[string]any{
			"requestId":     obs.RequestID,
			"method":        obs.Method,
			"url":           obs.URL,
			"host":          obs.Host,
			"path":          obs.Path,
			"query":         obs.Query,
			"contentType":   obs.ContentType,
			"blockedTarget": obs.BlockedTarget,
			"target":        obs.Target,
			"redirects":     obs.Redirects,
			"windowTargets": obs.WindowTargets,
		},
	}

	raw, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

func uniqueMatches(value string, tokens []string) []string {
	matches := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if strings.Contains(value, token) {
			matches = append(matches, token)
		}
	}

	return uniqueStrings(matches)
}

func uniqueStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}

	seen := make(map[string]struct{}, len(values))
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}

	return result
}

func collectText(obs Observation) []string {
	text := []string{
		obs.URL,
		obs.Host,
		obs.Path,
		obs.Query,
		obs.Body,
		obs.ContentType,
	}

	for key, value := range obs.Headers {
		text = append(text, key+"="+value)
	}

	return text
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}

	return ""
}

func stringPtr(value string) *string {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}

	return &value
}
