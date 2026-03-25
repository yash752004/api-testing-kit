# API Testing Kit - Project Plan

## 1. Project Summary

API Testing Kit is a developer-focused web application for creating API requests, executing them quickly, and inspecting responses in a clean, structured workspace.

This project should be built as a **portfolio-first live product**:
- polished enough to demonstrate publicly
- safe enough to run on the open internet
- scoped tightly enough to finish without turning into a full Postman clone

The core positioning is:

**A safe, developer-friendly API testing workspace where guests can explore the real app and signed-in users unlock full request execution.**

## Current Repository Baseline

- `apps/web` is the current frontend scaffold.
- `server` is the current Go backend scaffold.
- Only health endpoints exist today, so most backend behavior in this file is target-state planning.
- `schema.sql` is intentionally ahead of implementation.
- `goose` is the chosen migration tool once the database layer is wired in.
- Any doc that still describes Bun or `apps/api` as the active backend is stale and should be corrected.

---

## 2. Product Goals

### Primary Goals
- Demonstrate strong full-stack engineering and product thinking.
- Provide a real, usable API testing interface.
- Let visitors try the product immediately through restricted guest access on `/app`.
- Offer a signed-in mode for custom request execution with strong abuse controls.
- Present a clean request/response workflow that feels fast and professional.

### Secondary Goals
- Showcase system design, security awareness, and backend controls.
- Support reusable request collections and request history.
- Generate developer-friendly code snippets from requests.

### Non-Goals for V1
- Full team collaboration.
- Postman-level depth across every protocol and feature.
- GraphQL-specific advanced tooling.
- WebSocket testing.
- Multi-tenant enterprise permissions.
- Browser extension or desktop app.

---

## 3. Product Modes

The product must always be split into two modes.

### Guest App Mode
- No login required.
- Used for product exploration and portfolio demonstration.
- Only allowlisted guest endpoints or internal mock endpoints.
- No arbitrary external URL execution.

### Authenticated App Mode
- Login required.
- Supports user-created requests and custom target URLs.
- Protected by outbound request validation, quotas, abuse detection, and admin controls.

This split is not optional. The platform must never act like a public open proxy.

---

## 4. Target Users

### Primary Users
- Recruiters and clients reviewing the portfolio
- Developers exploring the guest-accessible app
- Individual developers who want a lightweight API testing workspace

### User Jobs
- Create and send HTTP requests quickly
- Inspect status, headers, and JSON response bodies
- Save useful requests into reusable collections
- Re-run recent requests
- Generate code snippets from working requests

---

## 5. Core Product Features

## MVP Features
- Request builder with `GET`, `POST`, `PUT`, `PATCH`, `DELETE`
- URL input with validation
- Query params editor
- Headers editor
- Auth selector for basic auth, bearer token, and no auth
- Body editor for JSON, raw text, and form-urlencoded
- Formatted response viewer
- Response metadata display: status, time, size
- Request history
- Saved collections
- Guest-safe templates/examples
- Generated code snippets for `curl`, `fetch`, and Python `requests`

## Phase 2 Features
- Environment variables
- Duplicate request
- Collection import/export
- Read-only shared collection links
- Basic API assertions such as status checks and simple JSON path checks

## Later Features
- Collection versioning
- Team workspaces
- More auth methods
- GraphQL support
- WebSocket testing
- Scheduled monitors

---

## 6. Core Pages

Detailed page structure lives in `ui-pages.md`, but the main routes are:

- `/` landing page
- `/app` shared guest/authenticated workspace
- `/templates` example collections
- `/features` product capability page
- `/docs` quick start and usage guide
- `/case-study` architecture and engineering page
- `/app/collections/[id]` collection detail page
- `/app/history` request history
- `/app/settings` settings and quotas
- `/admin/abuse` internal abuse monitoring page

---

## 7. User Experience Principles

- Keep request and response visible together at all times on desktop.
- Make guest access immediately useful without signup friction.
- Optimize for speed, clarity, and developer trust.
- Prefer strong defaults over excessive configuration.
- Make blocked or limited actions explicit with clear messaging.
- The guest experience on `/app` should feel real, but remain intentionally constrained.

---

## 8. Technology Stack

## Frontend
- SvelteKit
- Tailwind CSS
- Svelte-compatible component primitives if they materially speed up implementation, such as `shadcn-svelte` or an equivalent library
- Monaco Editor or CodeMirror for request/response editing

## Backend
- Go HTTP API in `server`
- Start with `net/http` and internal packages; add a heavier framework only if it solves a real problem
- REST API architecture
- HTTP client for outbound API execution

## Database
- PostgreSQL

## Infrastructure
- Docker
- Docker Compose for local multi-service setup

## Suggested Supporting Tools
- Redis for rate limiting, quotas, and short-lived counters
- `goose` for SQL migrations
- Zod for frontend validation
- `pgx`, `sqlc`, or another Go-friendly database access layer
- Server-side sessions with HTTP-only cookies
- Structured logging library

### Recommended Approach
- Use PostgreSQL for persistent product data.
- Use `goose` for versioned SQL migrations and treat `schema.sql` as the draft to convert into the initial migration set.
- Use Go for auth, request execution, and admin APIs so the docs match the current repo shape.
- Use Redis for rate limiting and abuse counters if available.
- If Redis is intentionally skipped for simplicity, document the limitation clearly and keep rate limiting conservative.

---

## 9. High-Level Architecture

### Frontend Responsibilities
- Render public marketing pages and the shared `/app` workspace
- Provide the request builder and response viewer UX
- Manage collections, history, and settings UI
- Display quota errors, blocked states, and validation feedback clearly

### Backend Responsibilities
- Handle authentication
- Store collections, requests, and history metadata
- Validate and execute outbound HTTP requests
- Enforce request validation and rate limits
- Log suspicious behavior
- Expose admin abuse insights

### Data Flow
1. User configures a request in the UI.
2. Frontend sends sanitized request payload to backend.
3. Backend validates user, quota, target, payload size, and destination safety.
4. Backend executes outbound request if allowed.
5. Backend trims or blocks oversized results if needed.
6. Backend returns structured response metadata and body preview.
7. Frontend renders formatted response, raw response, headers, and generated snippets.

---

## 10. Backend Modules

The backend should be split into clear modules:

- `auth`: signup, login, session management
- `users`: profile and account settings
- `collections`: collection CRUD
- `requests`: saved request definitions
- `history`: executed request records
- `runner`: outbound request validation and execution
- `rate-limit`: quota and abuse counters
- `abuse`: flagged events, denylist logic, admin review
- `templates`: guest-safe collections and sample requests
- `snippets`: code generation from request configuration

This separation matters because request execution and abuse protection are core system responsibilities, not helper logic.

---

## 11. Suggested Data Model

The exact schema can evolve, but the core entities should be:

- `users`
- `sessions`
- `collections`
- `saved_requests`
- `request_history`
- `request_templates`
- `usage_events`
- `abuse_events`
- `blocked_targets`

### Entity Intent
- `collections`: logical grouping of saved requests
- `saved_requests`: reusable request definitions
- `request_history`: executed runs with snapshots and metadata
- `usage_events`: counters and quota-related actions
- `abuse_events`: suspicious attempts and system blocks
- `blocked_targets`: denylisted domains, IPs, or patterns

Avoid over-modeling early. V1 should store only what is needed for product behavior, safety, and debugging.

---

## 12. Security and Abuse Prevention

This project has an unusually important security constraint:

**If outbound request execution is careless, the platform becomes a spam tool, scanning tool, or SSRF relay.**

So the security model must be part of the product plan, not an afterthought.

### Mandatory Rules
- Guest access on `/app` cannot send arbitrary external requests.
- Authenticated request execution must validate the destination before sending.
- Private and local network targets must be blocked.
- Cloud metadata targets must be blocked.
- Redirects must be limited and re-validated.
- Request size, response size, timeout, and concurrency must be capped.
- All suspicious attempts must be logged.

### Outbound Validation Rules
- Block `localhost`, `127.0.0.1`, `::1`
- Block RFC1918 private ranges
- Block link-local addresses
- Block metadata IPs like `169.254.169.254`
- Resolve DNS and validate resulting IPs
- Re-validate after redirects
- Restrict ports to safe defaults like `80` and `443` for launch
- Deny unsupported protocols

### Guest Access Rules
- Allow only curated guest endpoints
- No custom domains
- No webhook targets
- No file forwarding
- No large payloads
- No long-running requests

---

## 13. Rate Limiting Strategy

Rate limiting should exist at multiple levels.

### Guest App Limits
- 10 request executions per 10 minutes per IP
- 30 request executions per day per IP
- 1 active request at a time per IP
- 64 KB max request body
- 512 KB max response preview
- 10 second timeout

### Authenticated User Limits
- 60 request executions per hour per user
- 200 request executions per day per user
- 5 active concurrent requests per user
- 256 KB max request body
- 1 MB max response processing limit
- 15 second timeout
- 3 redirect max

### Global Protection
- Burst controls by IP
- Temporary cooldowns on spikes
- Domain-level throttling for repeated traffic
- Automatic account review or suspension on suspicious patterns

### Signup Safety
- CAPTCHA or Turnstile
- Email verification before custom URL execution
- Admin review tooling for abuse cases

---

## 14. Do's

- Do build the guest-accessible `/app` workspace first because it is the main portfolio surface.
- Do keep the UI centered around a strong request/response split workspace.
- Do ship a few polished guest-safe collections so the site feels alive immediately.
- Do keep the authenticated app simple in V1: single-user, no team complexity.
- Do treat security and abuse prevention as core product features.
- Do log blocked requests and suspicious patterns for admin visibility.
- Do favor clean, fast workflows over feature bloat.
- Do document intentional limitations openly so the product feels trustworthy.
- Do keep architecture modular so security-sensitive logic remains isolated.

---

## 15. Don'ts

- Don't expose arbitrary outbound request execution to anonymous users.
- Don't support too many advanced protocols in V1.
- Don't build a full Postman replacement before validating the core experience.
- Don't mix request execution logic directly into unrelated route handlers.
- Don't store more response data than needed for usability and debugging.
- Don't allow redirects to silently bypass validation.
- Don't treat rate limiting as optional.
- Don't overbuild multi-user or SaaS billing features before the core product is solid.
- Don't make the marketing site text-heavy and generic.

---

## 16. Development Phases

## Phase 1: Public-Facing Foundation
- Landing page
- Shared `/app` workspace
- Templates page
- Core request builder
- Core response viewer
- Guest-safe execution flow
- Basic public rate limiting

**Goal:** a live public site that feels complete and safe enough to share.

## Phase 2: Authenticated Product Layer
- Signup/login
- Authenticated request execution
- Collections
- History
- Code snippet generation
- Settings and quotas

**Goal:** a real signed-in product that supports individual developer workflows.

## Phase 3: Safety and Admin Depth
- Abuse monitor page
- Audit logging
- Denylist management
- Domain throttling
- Account review and suspension controls

**Goal:** operational visibility and better trustworthiness for live hosting.

## Phase 4: Productivity Enhancements
- Environment variables
- Import/export
- Shared read-only collections
- Basic assertions/tests

**Goal:** deeper utility without losing scope discipline.

## Phase 5: Optional Monetization Groundwork
- Plan and entitlement model
- Billing webhook skeleton
- Upgrade UI only if monetization becomes part of the project scope

**Goal:** leave room for pricing later without turning the MVP into a billing project.

---

## 17. MVP Success Criteria

The MVP is successful if:

- A visitor can land on the site and understand the product in under 30 seconds.
- A visitor can run example requests in guest mode on `/app` without confusion.
- A signed-in user can create, save, and re-run custom requests.
- The request/response UI feels fast and clean.
- The backend does not behave like an open proxy.
- Abuse controls are good enough for public hosting.
- The project demonstrates engineering maturity, not just UI polish.

---

## 18. Implementation Priorities

When tradeoffs appear, use this order:

1. Safety of outbound execution
2. Clean request/response UX
3. Guest app quality
4. Signed-in individual workflow
5. Productivity enhancements
6. Advanced feature depth

If a feature threatens delivery speed or security posture, it should be deferred.

---

## 19. Documentation Requirements

This project should maintain the following docs:

- `basic.md` for high-level project summary
- `ui-pages.md` for route and layout planning
- `design-guideline.md` for visual system and UI rules
- `plan.md` for source-of-truth execution plan
- `tasks.md` for actionable implementation work items
- `apis.md` for backend endpoints and contracts, with every route clearly labeled as implemented or planned

---

## 20. Final Direction

This project should not be framed as "yet another API client."

It should be framed as:

**A modern API testing workspace designed for fast demonstration, strong UX, and safe live deployment.**

That angle makes the project stronger as a portfolio piece because it shows:
- product thinking
- frontend craft
- backend architecture
- security awareness
- sensible scoping
