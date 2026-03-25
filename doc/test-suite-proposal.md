# API Testing Kit - Test Suite Proposal

## 1. Purpose

This document defines the testing approach for the current repo shape:
- `apps/web` is a SvelteKit frontend
- `server` is a Go backend
- `schema.sql` is the database contract draft

The project has a few high-risk areas:
- guest and authenticated behavior share the same `/app` surface
- the backend executes outbound HTTP requests
- SSRF prevention and abuse controls are part of core correctness
- the schema is broader than the code currently implemented

Because of that, the test plan should not rely only on browser smoke tests or only on backend unit tests.

## 2. Testing Philosophy

The right balance is:
- many fast unit tests
- a solid layer of integration tests around the Go server and database behavior
- a small set of high-value browser E2E flows
- dedicated security regression tests for outbound request safety

Avoid a giant E2E-first suite.

## 3. Recommended Test Layout

Use the current repo shape instead of planning around a non-existent `apps/api` package.

```text
apps/
  web/
    src/
    tests/
      unit/
      integration/
      e2e/
server/
  cmd/api/
  internal/
    ...
    *_test.go
  tests/
    integration/
      contracts/
      quotas/
      security/
test/
  fixtures/
  helpers/
  seed/
```

### Purpose Of Each Area

- `apps/web/tests/unit`: component behavior and pure UI state
- `apps/web/tests/integration`: route/load/action flows with mocked backend responses
- `apps/web/tests/e2e`: full browser journeys
- `server/internal/**/*_test.go`: package-level unit and service tests
- `server/tests/integration`: database-backed and HTTP-level server tests
- `test/helpers`: shared builders, DB setup, fake outbound targets, and seed helpers

## 4. Recommended Tooling

### Frontend
- `Vitest`
- `@testing-library/svelte`
- `Playwright`

### Backend
- Go `testing`
- `httptest`
- optional `testify` if it improves readability

### Database / Infra
- PostgreSQL test database
- `goose` for applying schema migrations in development and integration tests
- optional Redis test instance if rate limiting needs it
- optional `testcontainers` if isolated DB/Redis setup becomes painful

### Important Clarification

Do not plan backend testing around `Vitest` or `supertest`. The backend in this repo is Go, so the default server test runner should be `go test`.

## 5. Test Types

### A. Unit Tests

These should be the fastest and most common tests.

High-value unit targets:
- URL parsing and validation helpers
- request payload normalization
- auth/session helpers
- response formatting helpers
- guest lock-state helpers
- quota calculation helpers
- SSRF validation helpers

### B. Integration Tests

These should verify real collaboration across modules and the database boundary.

High-value integration targets:
- auth and session creation
- template loading
- request execution pipeline with a fake outbound HTTP target
- request history persistence
- collection CRUD
- abuse event creation
- rate-limit enforcement

### C. End-To-End Tests

Keep these few and high signal.

Core flows:
1. Guest opens `/app`, loads a template, runs an allowlisted request, and sees a formatted response.
2. Guest tries a locked action and sees the correct sign-in gating.
3. Signed-in user creates a request, sends it, saves it, and sees it in history.
4. Signed-in user cannot target a blocked local/private address.

## 6. Security Regression Suite

This project needs a dedicated security-focused suite.

Suggested path:

```text
server/tests/integration/security/
```

Cases to include:
- block `localhost`
- block `127.0.0.1`
- block `::1`
- block RFC1918 private ranges
- block `169.254.169.254`
- block unsupported protocols such as `file://`
- block redirect from safe host to blocked host
- block guest execution of arbitrary external URLs
- enforce timeout ceilings
- enforce request and response truncation ceilings

These are core product requirements, not edge cases.

## 7. Quota And Rate Limit Suite

Suggested path:

```text
server/tests/integration/quotas/
```

Cases to include:
- guest per-IP request limit
- guest daily limit
- authenticated per-user hourly limit
- authenticated concurrent request limit
- burst protection by IP
- abuse cooldown after repeated blocked attempts

## 8. Database And Schema Checks

Suggested focus:
- `goose up` applies cleanly on an empty PostgreSQL database
- `goose down` / rollback paths are verified where practical
- `schema.sql` stays aligned with the canonical `goose` migration history if it remains in the repo
- indexes and constraints behave as expected
- invalid HTTP URLs fail schema constraints where applicable
- `updated_at` triggers fire correctly
- template seed data loads cleanly once seeds exist

Treat `goose` migrations as the canonical migration path. If the draft `schema.sql` remains in the repo, keep it aligned with the migration history or clearly mark it as a generated reference.

## 9. API Contract Tests

Suggested path:

```text
server/tests/integration/contracts/
```

Verify:
- health response shape
- template list payload shape
- request execution response shape
- history payload shape
- collection CRUD payload shape

This becomes more important once the frontend and backend evolve in parallel.

## 10. Frontend Component Strategy

Do not over-test every visual component.

Test components where behavior matters:
- method selector
- auth scheme switcher
- request tab switching
- guest lock overlays
- response tabs
- history item selection
- collection list interactions

Avoid:
- snapshots for every card
- testing static marketing copy
- asserting Tailwind classes instead of behavior

## 11. Optional Monetization Tests

Billing and entitlements are later-phase work.

If that track starts, add service-level and contract tests then. Do not let billing become a prerequisite for the core runner, guest restrictions, or history flows.

## 12. Suggested Initial Backlog

Write these first:
- schema load test
- health endpoint contract test
- guest template fetch test
- request runner happy-path test
- private IP block test
- guest arbitrary URL denial test
- request history persistence test
- request builder component tests
- response viewer component tests
- E2E guest flow

## 13. Recommended Command Model

Keep the command model aligned to the current repo:

```bash
bun run test
bun run test:web:unit
bun run test:web:e2e
bun run test:server
```

Later, if needed, add narrower Go package or integration-specific commands without replacing the simple baseline.

## 14. Final Recommendation

The best testing stack for this repo is:
- `Vitest` plus `Playwright` for `apps/web`
- `go test` for `server`
- real PostgreSQL-backed integration checks for risky backend paths
- dedicated SSRF and quota regression suites

If those areas are covered well, the suite will give real confidence instead of shallow green checkmarks.
