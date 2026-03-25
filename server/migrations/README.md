This directory is the canonical goose migration history for the Go API.

The first migration intentionally captures the core identity and template tables from `schema.sql` so PostgreSQL-backed work can start without landing the entire future-state schema in one step.

Apply migrations with:

```bash
go run github.com/pressly/goose/v3/cmd/goose@v3.24.1 -dir ./server/migrations postgres "$DATABASE_URL" up
```
