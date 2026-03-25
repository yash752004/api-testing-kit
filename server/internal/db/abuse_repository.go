package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"api-testing-kit/server/internal/abuse"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AbuseRepository struct {
	pool *pgxpool.Pool
}

func NewAbuseRepository(pool *pgxpool.Pool) *AbuseRepository {
	return &AbuseRepository{pool: pool}
}

func (r *AbuseRepository) Create(ctx context.Context, event abuse.Event) (abuse.Event, error) {
	row := r.pool.QueryRow(ctx, `
		INSERT INTO abuse_events (
			user_id,
			session_id,
			request_run_id,
			severity,
			category,
			source_ip,
			target,
			rule_key,
			action_taken,
			details
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, user_id, session_id, request_run_id, severity, category, source_ip, target, rule_key, action_taken, details, created_at
	`,
		event.UserID,
		event.SessionID,
		event.RequestID,
		event.Severity,
		event.Category,
		event.SourceIP,
		event.Target,
		event.RuleKey,
		event.ActionTaken,
		event.Details,
	)

	var created abuse.Event
	if err := scanAbuseEvent(row.Scan, &created); err != nil {
		return abuse.Event{}, err
	}

	return created, nil
}

func (r *AbuseRepository) ListRecent(ctx context.Context, filter abuse.RecentFilter) ([]abuse.Event, error) {
	query, args := buildAbuseRecentQuery(filter)
	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]abuse.Event, 0)
	for rows.Next() {
		var item abuse.Event
		if err := scanAbuseEvent(rows.Scan, &item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r *AbuseRepository) SummarizeByCategory(ctx context.Context, filter abuse.SummaryFilter) ([]abuse.SummaryRow, error) {
	query, args := buildAbuseSummaryQuery(filter)
	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]abuse.SummaryRow, 0)
	for rows.Next() {
		var item abuse.SummaryRow
		if err := rows.Scan(
			&item.Severity,
			&item.Category,
			&item.ActionTaken,
			&item.Count,
			&item.LastCreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func buildAbuseRecentQuery(filter abuse.RecentFilter) (string, []any) {
	var clauses []string
	args := make([]any, 0, 10)

	addClause := func(expr string, value any) {
		args = append(args, value)
		clauses = append(clauses, fmt.Sprintf(expr, len(args)))
	}

	if filter.UserID != nil {
		addClause("user_id = $%d", *filter.UserID)
	}
	if filter.SessionID != nil {
		addClause("session_id = $%d", *filter.SessionID)
	}
	if filter.RequestRunID != nil {
		addClause("request_run_id = $%d", *filter.RequestRunID)
	}
	if filter.Severity != "" {
		addClause("severity = $%d", filter.Severity)
	}
	if filter.Category != "" {
		addClause("category = $%d", filter.Category)
	}
	if filter.SourceIP != nil {
		addClause("source_ip = $%d", *filter.SourceIP)
	}
	if filter.Target != nil {
		addClause("target = $%d", *filter.Target)
	}
	if filter.RuleKey != nil {
		addClause("rule_key = $%d", *filter.RuleKey)
	}
	if filter.ActionTaken != "" {
		addClause("action_taken = $%d", filter.ActionTaken)
	}
	if filter.CreatedAfter != nil {
		addClause("created_at >= $%d", *filter.CreatedAfter)
	}
	if filter.CreatedBefore != nil {
		addClause("created_at <= $%d", *filter.CreatedBefore)
	}

	limit := normalizeLimit(filter.Limit, 100, 500)
	args = append(args, limit)

	var builder strings.Builder
	builder.WriteString(`
		SELECT id, user_id, session_id, request_run_id, severity, category, source_ip, target, rule_key, action_taken, details, created_at
		FROM abuse_events
	`)
	if len(clauses) > 0 {
		builder.WriteString(" WHERE ")
		builder.WriteString(strings.Join(clauses, " AND "))
	}
	builder.WriteString(fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d", len(args)))

	return builder.String(), args
}

func scanAbuseEvent(scan func(dest ...any) error, item *abuse.Event) error {
	var userID sql.NullString
	var sessionID sql.NullString
	var requestID sql.NullString
	var sourceIP sql.NullString
	var target sql.NullString
	var ruleKey sql.NullString

	if err := scan(
		&item.ID,
		&userID,
		&sessionID,
		&requestID,
		&item.Severity,
		&item.Category,
		&sourceIP,
		&target,
		&ruleKey,
		&item.ActionTaken,
		&item.Details,
		&item.CreatedAt,
	); err != nil {
		return err
	}
	if userID.Valid {
		value := userID.String
		item.UserID = &value
	}
	if sessionID.Valid {
		value := sessionID.String
		item.SessionID = &value
	}
	if requestID.Valid {
		value := requestID.String
		item.RequestID = &value
	}
	if sourceIP.Valid {
		value := sourceIP.String
		item.SourceIP = &value
	}
	if target.Valid {
		value := target.String
		item.Target = &value
	}
	if ruleKey.Valid {
		item.RuleKey = ruleKey.String
	}
	return nil
}

func buildAbuseSummaryQuery(filter abuse.SummaryFilter) (string, []any) {
	var clauses []string
	args := make([]any, 0, 8)

	addClause := func(expr string, value any) {
		args = append(args, value)
		clauses = append(clauses, fmt.Sprintf(expr, len(args)))
	}

	if filter.UserID != nil {
		addClause("user_id = $%d", *filter.UserID)
	}
	if filter.SessionID != nil {
		addClause("session_id = $%d", *filter.SessionID)
	}
	if filter.RequestRunID != nil {
		addClause("request_run_id = $%d", *filter.RequestRunID)
	}
	if filter.Severity != "" {
		addClause("severity = $%d", filter.Severity)
	}
	if filter.Category != "" {
		addClause("category = $%d", filter.Category)
	}
	if filter.SourceIP != nil {
		addClause("source_ip = $%d", *filter.SourceIP)
	}
	if filter.CreatedAfter != nil {
		addClause("created_at >= $%d", *filter.CreatedAfter)
	}
	if filter.CreatedBefore != nil {
		addClause("created_at <= $%d", *filter.CreatedBefore)
	}

	limit := normalizeLimit(filter.Limit, 50, 200)
	args = append(args, limit)

	var builder strings.Builder
	builder.WriteString(`
		SELECT severity, category, action_taken, COUNT(*) AS count, MAX(created_at) AS last_created_at
		FROM abuse_events
	`)
	if len(clauses) > 0 {
		builder.WriteString(" WHERE ")
		builder.WriteString(strings.Join(clauses, " AND "))
	}
	builder.WriteString(`
		GROUP BY severity, category, action_taken
		ORDER BY MAX(created_at) DESC, severity ASC, category ASC, action_taken ASC
		LIMIT $`)
	builder.WriteString(fmt.Sprint(len(args)))

	return builder.String(), args
}
