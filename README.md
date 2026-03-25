# API Testing Kit

API Testing Kit is a monorepo for a developer-focused API testing product.

## Current Repository State

- `apps/web` is the SvelteKit frontend scaffold and still contains mostly starter routes.
- `server` is the Go HTTP scaffold. Only health endpoints are implemented today.
- `schema.sql` is a future-state PostgreSQL schema draft and is ahead of the current backend implementation.
- `goose` is the chosen database migration tool once persistence work starts. The current `schema.sql` should be converted into the initial migration set.
- `doc/` contains the target product, UX, and execution docs. Those docs should explicitly separate implemented behavior from planned behavior.

## Repository Structure

```text
api-testing-kit/
  apps/
    web/
  server/
  doc/
  schema.sql
  package.json
  README.md
```

## Packages

### Frontend
- Path: `apps/web`
- Stack: SvelteKit, TypeScript, Tailwind CSS
- Current status: starter scaffold that still needs to be aligned to the `/app` product direction

### Backend
- Path: `server`
- Stack: Go
- Current status: minimal HTTP scaffold with health endpoints

### Documentation
- Path: `doc`
- Includes:
  - `basic.md` for the project summary and locked decisions
  - `plan.md` for the target product direction and phased scope
  - `ui-pages.md` for route and workspace planning
  - `design-guideline.md` for the visual system
  - `tasks.md` for implementation sequencing
  - `test-suite-proposal.md` for test planning
  - `apis.md` for implemented endpoints and planned API contracts

## Root Commands

### Frontend

```bash
bun run dev:web
bun run build:web
bun run preview:web
bun run check:web
bun run test:web
```

### Backend

```bash
bun run dev:server
bun run test:server
```

### Combined

```bash
bun run test
```

## Environment

Copy values from `.env.example` before local development.

Current variables:

```bash
PUBLIC_API_BASE_URL=http://localhost:8080
API_PORT=8080
```

## Backend Endpoints

The current Go scaffold exposes:

- `/`
- `/healthz`
- `/api/v1/health`

## Source Of Truth

- Current implementation state: this README plus the code under `apps/web`, `server`, and `schema.sql`
- Target product and UX direction: `doc/plan.md`, `doc/ui-pages.md`, and `doc/design-guideline.md`
- Execution backlog: `doc/tasks.md`
- Backend contract decisions: `doc/apis.md`
