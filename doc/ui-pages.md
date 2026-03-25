# API Testing Kit - UI Pages

## Product Modes

The product should use one shared app route, `/app`, with two different capability levels.

### 1. Guest App Mode
- Purpose: Let visitors explore the real workspace without signup friction.
- Access: No login required.
- Route: `/app`
- Restriction: Guests cannot use arbitrary external URLs.
- Data Source: Only allowlisted sample APIs or internal mock endpoints controlled by us.
- Experience: Guests see the real app layout, but with locked advanced actions and upgrade prompts.

### 2. Authenticated App Mode
- Purpose: Let signed-in users create, save, and run their own requests.
- Access: Login required.
- Route: `/app`
- Restriction: Protected by strict outbound request validation, quotas, abuse detection, and logging.
- Experience: Same workspace, but with authenticated capabilities unlocked. If monetization is added later, plan entitlements can layer on top without changing the core app model.

This shared-route model is important. Visitors should experience the real product surface immediately, while sensitive functionality stays gated behind login and backend controls.

---

## Public Site Pages

## 1. Landing Page
**Route:** `/`

**Purpose**
- Introduce the product clearly.
- Show that this is an API testing and response inspection tool.
- Push visitors toward the live app at `/app`.

**Sections**
- Hero section with strong headline and short product summary.
- Visual product preview showing split request/response workspace.
- Feature highlights.
- Example use cases.
- CTA to open the app.
- Technical credibility section.
- Footer with links to app, docs, source, and case study.

**Layout Structure**
- Top navigation bar.
- Hero with left text block and right product screenshot/mockup.
- Feature grid section.
- Templates/examples section.
- Architecture/tech stack section.
- CTA banner.
- Footer.

---

## 2. App Workspace
**Route:** `/app`

**Purpose**
- Serve as the single working surface for both guests and logged-in users.
- Demonstrate the product immediately while preserving security and upgrade paths.

**Guest App Behavior**
- Guests can browse predefined sample collections.
- Guests can inspect request configuration and response structure.
- Guests can edit only safe fields explicitly marked as overridable inside curated templates.
- Guests can run only allowlisted endpoints.
- Guests must not be able to replace the template target with an arbitrary domain or URL.
- Guests cannot save permanent collections or history.
- Guests cannot use secret environment variables.
- Guests should see clear upgrade or sign-in prompts on locked features.

**Authenticated App Behavior**
- Signed-in users can create requests with multiple HTTP methods.
- Signed-in users can configure params, headers, auth, and body.
- Signed-in users can run validated custom URLs.
- Signed-in users can save requests to collections.
- Signed-in users can view history and reuse requests.
- Signed-in users can use account-level quotas. If monetization is added later, entitlements can layer on top, but the MVP should not depend on billing.

**Main UI**
- Top bar with logo, templates, docs, sign in/account menu, and an upgrade CTA only if monetization exists.
- Left sidebar for collections, examples, and history access.
- Main request panel for method, URL, params, headers, auth, and body.
- Right response panel for pretty JSON, raw response, headers, and request metadata.
- Bottom utility drawer for generated code snippets. Test results belong to a later assertions phase, not the base MVP workspace.

**Layout Structure**
- Header.
- Three-column workspace:
  - Left: collections/examples/history entry points
  - Center: request editor
  - Right: response viewer
- Bottom drawer for secondary tooling.

**Default Guest Content**
- Preload examples such as:
  - JSONPlaceholder
  - GitHub public API
  - Weather demo endpoint
  - Auth flow mock example

---

## 3. Templates Page
**Route:** `/templates`

**Purpose**
- Show ready-made API request collections.
- Help visitors understand real-world use cases quickly.

**Sections**
- Search and filter bar.
- Template cards grouped by category.
- Collection preview drawer or modal.
- CTA to open any template in `/app`.

**Suggested Categories**
- REST basics
- Authentication flows
- CRUD examples
- Pagination examples
- Webhooks
- Error handling

---

## 4. Features Page
**Route:** `/features`

**Purpose**
- Explain product capabilities in a structured way.
- Make the site feel like a real product, not just a mockup.

**Sections**
- Request builder features.
- Response inspection features.
- Collection management.
- Code snippet generation.
- History and saved requests.
- Guest limitations and authenticated unlocks.
- Safety and abuse prevention model.

---

## 5. Docs / Quick Start Page
**Route:** `/docs`

**Purpose**
- Explain how to use the tool quickly.
- Reduce friction for first-time visitors.

**Sections**
- Quick start.
- How to create a request.
- How to inspect responses.
- How collections work.
- What guests can do in `/app`.
- What unlocks after sign-in.

---

## 6. Case Study / Architecture Page
**Route:** `/case-study`

**Purpose**
- Demonstrate engineering thinking to employers, clients, and developers.
- Explain architecture, constraints, and design choices.

**Sections**
- Problem statement.
- Product goals.
- Frontend architecture.
- Backend architecture.
- Request execution pipeline.
- Guest gating model on `/app`.
- Abuse prevention strategy.
- Rate limiting strategy.
- Deployment overview.

---

## Authenticated App Pages

## 7. Collection Detail Page
**Route:** `/app/collections/[id]`

**Purpose**
- Organize related requests into reusable collections.

**Features**
- Folder-like grouping.
- Request list.
- Rename, duplicate, delete.
- Share read-only collection link if enabled.

**Layout Structure**
- Collection header with title and actions.
- Left request tree.
- Center request editor.
- Right response/test panel.

**Access**
- Login required.

---

## 8. History Page
**Route:** `/app/history`

**Purpose**
- Show previously executed requests.

**Features**
- Filter by status, domain, method, and date.
- Re-run request.
- Save history item into a collection.
- Inspect response snapshot.

**Access**
- Login required.

---

## 9. Settings Page
**Route:** `/app/settings`

**Purpose**
- Manage account-level preferences and execution rules.

**Sections**
- Profile.
- API environments and variables.
- Usage and quotas.
- Security notices.
- Export/import settings.

**Access**
- Login required.

---

## 10. Admin Abuse Monitor
**Route:** `/admin/abuse`

**Purpose**
- Monitor suspicious behavior and enforce safety controls.

**Features**
- Request volume dashboard.
- Flagged IPs and users.
- Recent blocked requests.
- Domain and IP denylist management.
- User suspension controls.
- Audit log view.

**Access**
- Internal/admin only.

---

## Shared UI Blocks

These components should be reused across guest and authenticated app states:

- Method selector with clear color coding.
- URL input with validation states.
- Params table editor.
- Headers table editor.
- Auth selector.
- Body editor with JSON formatting support.
- Response status/time/size badges.
- Pretty/raw/headers response tabs.
- Code snippet panel.
- Locked-feature callouts for guests.
- Empty, loading, error, and blocked-state components.

---

## Page Layout Rules

### Marketing Pages
- Wide layout with clear section separation.
- Strong top-level CTA toward `/app`.
- Use large product visuals, not text-heavy blocks.

### App Workspace Pages
- Desktop-first split layout.
- Keep request and response visible at the same time.
- Sidebar should remain persistent on desktop.
- On mobile/tablet, collapse sidebar and stack panels vertically.
- Locked guest actions should be visible but clearly disabled or upsold.

### Core Workspace Pattern
- Header
- Sidebar
- Request Builder
- Response Viewer
- Utility Drawer

This request/response split is the central visual identity of the product.

---

## Guest App Safety Rules

Guest access to `/app` must be intentionally limited:

- No arbitrary external URL entry.
- Only allowlisted demo domains or internal mock endpoints.
- No private network targets.
- No custom webhook targets.
- No file upload forwarding.
- No large request bodies.
- No long-running requests.
- No persistence of sensitive user data.

If a visitor wants full request control, require account creation and move them into authenticated mode.

---

## Outbound Request Validation Rules

These rules apply to any user-generated outbound request in authenticated mode:

- Block requests to `localhost`, `127.0.0.1`, `::1`.
- Block RFC1918 private ranges:
  - `10.0.0.0/8`
  - `172.16.0.0/12`
  - `192.168.0.0/16`
- Block link-local and local network ranges.
- Block cloud metadata targets such as `169.254.169.254`.
- Resolve DNS and validate final IP before sending.
- Re-check on redirect hops.
- Restrict allowed ports to standard safe ports like `80` and `443` at launch.
- Deny unsupported protocols.
- Deny redirect chains beyond a small limit.

This is required to reduce SSRF and infrastructure abuse risk.

---

## Rate Limiting Rules

Rate limiting must exist at multiple layers:

### Guest App Limits
- Max 10 request executions per 10 minutes per IP.
- Max 30 request executions per day per IP.
- Only allowlisted guest endpoints.
- Max request body size: 64 KB.
- Max response size shown: 512 KB.
- Max request timeout: 10 seconds.
- Max concurrency: 1 active request per IP.

### Authenticated User Limits
- Max 60 request executions per hour per user.
- Max 200 request executions per day per user.
- Max 5 active concurrent requests per user.
- Max request body size: 256 KB.
- Max response size processed: 1 MB.
- Max request timeout: 15 seconds.
- Max redirect count: 3.

### IP-Based Global Protection
- Max burst traffic per IP even across multiple accounts.
- Temporary cooldown on repeated failures or spikes.
- Automatic block on suspicious volume patterns.

### Domain-Level Protection
- Rate limit repeated requests to the same target domain.
- Slow down or block domains with repeated abuse signals.
- Maintain a denylist and manual review flow.

### Account Safety Rules
- CAPTCHA or Turnstile on signup.
- Email verification before enabling custom URL execution.
- Automatic suspension for suspicious usage patterns.
- Admin review tools for restoring or permanently blocking accounts.

---

## Abuse Signals To Monitor

- High request volume in short time windows.
- Repeated requests to many unrelated domains.
- Repeated failed DNS lookups.
- Repeated attempts to hit blocked IP ranges.
- Repeated timeout-heavy traffic.
- Pattern matching consistent with scanning, scraping, or spam relay behavior.

Any such signal should trigger logging, throttling, and optional suspension.

---

## MVP Priority Order

### Phase 1
- Landing page
- Shared `/app` workspace
- Guest restrictions and locked states
- Core request builder
- Core response viewer
- Demo templates

### Phase 2
- Login and authenticated full app behavior
- Saved collections
- History
- Code snippet generation

### Phase 3
- Abuse monitor admin page
- Domain/IP denylist tools
- Quota and usage dashboard
- Shared collection links

---

## Final Product Positioning

This should be presented publicly as:

**A safe, developer-friendly API testing workspace where anyone can explore the real app, and signed-in users unlock full request execution.**

That framing is important because it keeps the product feeling real while making guest limitations intentional rather than incomplete.
