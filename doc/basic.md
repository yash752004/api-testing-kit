# API Testing Kit

## Product Summary
- Project Type: Developer-focused API testing workspace
- Product Goal: Let users compose HTTP requests quickly, execute them safely, and inspect structured responses without turning the product into an open proxy
- Core Product Split: guest exploration on `/app` plus authenticated custom execution with backend safety controls

## Current Repository Baseline
- Frontend scaffold: `apps/web` (SvelteKit)
- Backend scaffold: `server` (Go HTTP server)
- Database contract: `schema.sql` (ahead of current implementation)
- Current implemented backend surface: health endpoints only

## MVP Target
- Shared `/app` request/response workspace
- HTTP methods: `GET`, `POST`, `PUT`, `PATCH`, `DELETE`
- Request configuration: URL, query params, headers, auth, and body
- Response inspection: status, headers, timing, size, and formatted body
- Guest-safe templates/examples
- Authenticated collections and request history after the core runner is stable

## Technology Decisions
- Frontend: SvelteKit, TypeScript, Tailwind CSS
- UI Primitives: Svelte-compatible component primitives if they materially speed up implementation, such as `shadcn-svelte` or an equivalent library
- Backend: Go HTTP API in `server`
- Database: PostgreSQL
- Migrations: `goose`
- Infrastructure: Docker / Docker Compose
- Testing: Vitest and Playwright for `apps/web`, `go test` for `server`

## Scope Boundaries
- Billing and entitlements are later-phase work, not an MVP blocker.
- `apis.md` must distinguish between implemented endpoints and planned endpoints.
- Any doc that still describes Bun or `apps/api` as the active backend should be treated as stale.
