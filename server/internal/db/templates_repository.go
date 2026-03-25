package db

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TemplateCollection struct {
	ID            string
	Slug          string
	Name          string
	Description   string
	Visibility    string
	Category      string
	FeaturedOrder int
	IsActive      bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type RequestTemplate struct {
	ID                   string
	TemplateCollectionID *string
	Slug                 string
	Name                 string
	Description          string
	Visibility           string
	Category             string
	RequestDefinition    json.RawMessage
	PreviewResponse      json.RawMessage
	SortOrder            int
	FeaturedOrder        int
	IsActive             bool
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type TemplateRepository struct {
	pool *pgxpool.Pool
}

func NewTemplateRepository(pool *pgxpool.Pool) *TemplateRepository {
	return &TemplateRepository{pool: pool}
}

func (r *TemplateRepository) ListActiveCollections(ctx context.Context) ([]TemplateCollection, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT
			id,
			slug,
			name,
			COALESCE(description, ''),
			visibility,
			COALESCE(category, ''),
			featured_order,
			is_active,
			created_at,
			updated_at
		FROM template_collections
		WHERE is_active = TRUE
		ORDER BY featured_order ASC, slug ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]TemplateCollection, 0)
	for rows.Next() {
		var item TemplateCollection
		if err := rows.Scan(
			&item.ID,
			&item.Slug,
			&item.Name,
			&item.Description,
			&item.Visibility,
			&item.Category,
			&item.FeaturedOrder,
			&item.IsActive,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, rows.Err()
}

func (r *TemplateRepository) ListActiveTemplates(ctx context.Context) ([]RequestTemplate, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT
			id,
			template_collection_id,
			slug,
			name,
			COALESCE(description, ''),
			visibility,
			COALESCE(category, ''),
			request_definition,
			preview_response,
			sort_order,
			featured_order,
			is_active,
			created_at,
			updated_at
		FROM request_templates
		WHERE is_active = TRUE
		ORDER BY featured_order ASC, sort_order ASC, slug ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]RequestTemplate, 0)
	for rows.Next() {
		var item RequestTemplate
		if err := rows.Scan(
			&item.ID,
			&item.TemplateCollectionID,
			&item.Slug,
			&item.Name,
			&item.Description,
			&item.Visibility,
			&item.Category,
			&item.RequestDefinition,
			&item.PreviewResponse,
			&item.SortOrder,
			&item.FeaturedOrder,
			&item.IsActive,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, rows.Err()
}
