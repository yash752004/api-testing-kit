# API Testing Kit - Design Guideline

## 1. Design Intent

The visual style for this project should be derived from the provided reference image:

- soft premium SaaS dashboard
- rounded, calm, high-trust surfaces
- light-first interface with warm neutral backgrounds
- green as the dominant brand/action color
- dense information presented in clean modular cards
- polished but not flashy

This style should be adapted for an **API testing product**, not copied literally as a project-management dashboard.

The result should feel like:

**A premium developer workspace with the warmth of a modern SaaS dashboard and the clarity of a technical tool.**

---

## 2. Visual Personality

The UI should communicate:

- trustworthy
- clean
- premium
- calm
- precise
- efficient

The interface should not feel:

- dark and hacker-themed
- overly corporate and cold
- too minimal to the point of emptiness
- colorful in a noisy way
- glassy, neon, or futuristic

---

## 3. Core Style Principles

### 1. Soft Container, Structured Interior
- The overall application should sit inside a large rounded shell.
- Individual content areas should be broken into smaller rounded cards.
- Cards should feel lightly elevated, not heavily shadowed.

### 2. Warm Light Theme
- Use soft ivory, stone, and white surfaces instead of pure white everywhere.
- Backgrounds should feel slightly warm and tactile.

### 3. Green-Driven Action System
- Primary actions and positive states should use rich green.
- Green should appear in buttons, active nav states, charts, badges, and highlights.
- Avoid introducing too many other strong colors.

### 4. Dense But Calm
- The layout can contain many modules, but spacing must remain generous enough to prevent stress.
- Use clear card boundaries and strong alignment to control density.

### 5. Real Product Feel
- Locked features, quotas, and app states should feel intentional.
- Guests should see the real product layout, not a fake simplified version.

---

## 4. Brand Direction For This Project

Since this is an API testing workspace, the visual language should translate the reference style into a developer-focused product:

- replace project dashboard cards with request, response, history, and collections modules
- keep the same rounded shell, card system, and soft green accent model
- maintain strong readability for code, JSON, headers, and form fields
- balance friendly SaaS styling with technical precision

This should look like a developer tool built by a product-minded team, not an internal admin panel.

---

## 5. Color System

## Primary Palette

Use a warm neutral foundation with green accents.

- `Canvas`: `#efeeea`
- `Shell`: `#f7f5f0`
- `Surface`: `#ffffff`
- `Surface Soft`: `#f5f3ed`
- `Border Soft`: `#e7e3d8`
- `Text Strong`: `#162117`
- `Text Body`: `#445046`
- `Text Muted`: `#7a847d`
- `Primary Green`: `#1f7a4d`
- `Primary Green Hover`: `#19663f`
- `Primary Green Soft`: `#dcefe3`
- `Primary Green Deep`: `#145336`

## Semantic Colors

- `Success`: `#2d8a57`
- `Warning`: `#f0b44c`
- `Danger`: `#e36d5d`
- `Info`: `#6f8ea3`

## Accent Usage Rules

- Green is the dominant accent.
- Yellow/orange is reserved for warning or pending states.
- Red is reserved for destructive or failed states.
- Blue should be used sparingly and only where status distinction is necessary.

Do not create rainbow-heavy dashboards.

---

## 6. Component Theme Tokens

Use a token system that works with Svelte component primitives. If `shadcn-svelte` is used, treat these as the base tokens. Do not hardcode custom hex values inside every component.

Suggested light theme token mapping:

```css
:root {
  --background: 42 24% 93%;
  --foreground: 135 18% 11%;

  --card: 0 0% 100%;
  --card-foreground: 135 18% 11%;

  --popover: 0 0% 100%;
  --popover-foreground: 135 18% 11%;

  --primary: 147 58% 30%;
  --primary-foreground: 0 0% 100%;

  --secondary: 42 27% 91%;
  --secondary-foreground: 135 16% 18%;

  --muted: 42 18% 90%;
  --muted-foreground: 135 8% 43%;

  --accent: 140 35% 89%;
  --accent-foreground: 140 32% 20%;

  --destructive: 8 68% 61%;
  --destructive-foreground: 0 0% 100%;

  --border: 42 18% 85%;
  --input: 42 18% 85%;
  --ring: 147 58% 30%;

  --radius: 1.25rem;

  --chart-1: 147 58% 30%;
  --chart-2: 147 36% 45%;
  --chart-3: 147 28% 58%;
  --chart-4: 42 20% 72%;
  --chart-5: 42 16% 82%;
}
```

### Additional Custom Tokens

Add project-specific variables in `src/routes/layout.css` or an equivalent global theme file:

```css
:root {
  --shell: #f7f5f0;
  --canvas: #efeeea;
  --panel-soft: #f5f3ed;
  --success: #2d8a57;
  --warning: #f0b44c;
  --danger: #e36d5d;
  --shadow-soft: 0 24px 60px rgba(21, 31, 23, 0.08);
  --shadow-card: 0 10px 30px rgba(21, 31, 23, 0.05);
}
```

---

## 7. Tailwind Theme Extensions

Extend Tailwind so the visual system is reusable and not recreated ad hoc.

Suggested additions:

```ts
extend: {
  colors: {
    shell: "var(--shell)",
    canvas: "var(--canvas)",
    "panel-soft": "var(--panel-soft)",
    success: "var(--success)",
    warning: "var(--warning)",
    danger: "var(--danger)",
  },
  boxShadow: {
    shell: "0 24px 60px rgba(21, 31, 23, 0.08)",
    card: "0 10px 30px rgba(21, 31, 23, 0.05)",
    float: "0 18px 40px rgba(21, 31, 23, 0.10)",
  },
  borderRadius: {
    xl2: "1.25rem",
    xl3: "1.75rem",
    xl4: "2rem",
  },
  backgroundImage: {
    "stripe-soft":
      "repeating-linear-gradient(135deg, rgba(22,33,23,0.12) 0, rgba(22,33,23,0.12) 2px, transparent 2px, transparent 8px)",
    "green-radial":
      "radial-gradient(circle at top left, rgba(48,130,84,0.24), transparent 48%)",
  },
}
```

---

## 8. Typography

The reference style uses rounded, friendly, product-oriented typography.

### Recommended Font Pairing

- UI Font: `Manrope`
- Code Font: `IBM Plex Mono`

### Usage Rules

- Use `Manrope` for navigation, headings, buttons, labels, and body text.
- Use `IBM Plex Mono` for URLs, status metadata, response size/time, header keys, JSON code, and snippet blocks.

### Type Scale

- Page Title: `text-4xl font-semibold tracking-tight`
- Section Title: `text-2xl font-semibold`
- Card Title: `text-base font-semibold`
- Body: `text-sm leading-6`
- Meta Text: `text-xs text-muted-foreground`
- Code Meta: `text-xs font-medium font-mono`

### Typography Rules

- Use tight tracking on headings.
- Keep body text neutral and readable.
- Do not use oversized hero typography inside the app workspace.
- Avoid thin fonts.

---

## 9. Surface System

The interface should be built in layers:

### Layer 1. Page Canvas
- Full-screen warm neutral background
- Very subtle gradient or blur is acceptable
- Never use a flat pure white full-screen background

### Layer 2. App Shell
- Large centered rounded container
- Background: shell/off-white
- Soft border and large shadow
- Padding: comfortable and generous

### Layer 3. Internal Panels
- Sidebar, header strip, request panel, response panel
- White or near-white cards
- Rounded corners
- Thin soft border

### Layer 4. Interactive Modules
- Inputs, tabs, metric cards, history items, code blocks
- More compact but still soft and rounded

---

## 10. Border Radius and Corners

Rounded corners are a major part of this style.

### Radius Rules

- Outer shell: `rounded-[32px]`
- Main header and large panels: `rounded-[24px]`
- Standard cards: `rounded-[20px]`
- Inputs and buttons: `rounded-full` or `rounded-[16px]`
- Small pills and badges: `rounded-full`

Avoid sharp corners entirely in the main UI.

---

## 11. Shadow and Border Treatment

Shadows should be soft and low contrast.

### Rules

- Prefer soft depth over heavy elevation
- Use shadows to separate large surfaces, not every tiny element
- Pair shadows with subtle borders

### Border Rules

- Border color should be warm and very light
- Use `border-border/80` or equivalent
- Avoid dark gray borders

### Example Style

- shell: strong soft shadow
- cards: minimal soft shadow
- inputs: mostly border-driven, little to no shadow

---

## 12. Layout Structure

## App-Level Shell

The main interface should mirror the reference image:

- centered shell container
- left vertical sidebar
- top toolbar
- modular content grid inside the main area

### Suggested Desktop Shell

- max width: `1440px`
- min height: `calc(100vh - 32px)`
- shell padding: `12px` to `16px`
- internal gap: `12px`

### Sidebar

- width: `240px` to `260px`
- persistent on desktop
- contains logo, primary nav, secondary nav, promotional/help card

### Main Area

- top toolbar card
- content region below
- request and response panels arranged as a split workspace

---

## 13. API Testing Workspace Translation

This project should adapt the dashboard language into an API workspace.

### Recommended `/app` Structure

- Sidebar
- Toolbar
- Request metrics row
- Main request/response split
- Lower utility row for snippets, history, and saved items

### Suggested Main Grid

- Left sidebar: collections, templates, recent requests
- Center panel: request builder
- Right panel: response viewer

### Optional Internal Arrangement

- Top row inside main area:
  - quick stats cards for requests sent, saved collections, quota remaining
- Middle row:
  - large request editor card
  - large response card
- Bottom row:
  - code snippets
  - request history
  - feature lock or upgrade card for guests

This preserves the premium dashboard feel while still being useful for API work.

---

## 14. Header / Toolbar Design

The toolbar should visually follow the reference image:

- rounded full-width card
- search or quick action field on the left
- compact utility icons in circular buttons
- profile/account menu on the right

### Toolbar Content For This Product

- collection search or request search
- environment switcher
- quota badge
- notifications or docs button
- account menu or sign-in CTA

### Toolbar Styling

- height around `72px`
- background: white
- rounded `24px`
- horizontal padding: `20px` to `24px`

---

## 15. Sidebar Design

The sidebar should be soft and quiet, like the reference image.

### Structure

- logo block at top
- primary nav group
- secondary nav group
- bottom support or upsell card

### Nav Item Style

- height: `40px` to `44px`
- icon on left
- label text
- optional badge on right
- default: muted text
- hover: slightly darker text, soft background
- active: green-filled pill or green left accent with tinted background

### Sidebar Components

- App
- Templates
- Collections
- History
- Usage
- Settings
- Help

---

## 16. Card System

Cards are the core building block.

### Standard Card

- white background
- rounded `20px`
- soft border
- subtle shadow
- padding `20px` to `24px`

### Highlight Card

- green or green-tinted background
- used sparingly for:
  - quota summary
  - usage summary
  - CTA blocks
  - featured template

### Soft Utility Card

- off-white background
- minimal shadow
- used for filters, empty states, guest locks

---

## 17. Button Styles

Buttons should look rounded, polished, and intentional.

### Primary Button

- filled green background
- white text
- rounded full pill
- medium weight
- subtle hover darken

Use for:
- Send Request
- Save Request
- Sign In
- Upgrade

### Secondary Outline Button

- white or shell background
- soft green or neutral border
- dark text
- rounded full pill

Use for:
- Import
- Duplicate
- Open Docs

### Ghost Button

- transparent background
- muted text
- hover with soft tint

Use for:
- icon actions
- subtle toolbar controls

### Destructive Button

- use only for delete/remove actions
- danger text or danger tint

---

## 18. Input and Form Controls

The reference image uses soft, rounded controls with plenty of breathing room.

### Inputs

- rounded full or `rounded-[16px]`
- white background
- soft border
- light placeholder
- no harsh inset shadows

### URL Input

- larger than standard fields
- include method selector attached or adjacent
- use monospace for the URL itself if readability benefits

### Table-Like Editors

For query params and headers:

- use compact rounded rows
- each row should feel like a mini card
- use alternating soft background on hover
- keep row borders subtle

### Auth Selector

- use segmented or tabbed control
- avoid native-looking select boxes when possible

---

## 19. Tabs and Segment Controls

Tabs are important for request and response flows.

### Use Cases

- Params / Headers / Auth / Body
- Pretty / Raw / Headers / Snippets

### Style

- rounded full segmented container
- active tab can be green tint or white with stronger border
- inactive tabs stay muted

The segmented control should feel premium and tactile, not browser-default.

---

## 20. Response Viewer Styling

The response viewer must remain technical and readable inside this soft visual system.

### Response Meta Row

At the top of the response card include:

- status badge
- response time
- response size
- content type

### Response Body

- use a white or soft off-white code surface
- monospace font
- syntax highlighting should remain restrained
- avoid black editor themes in the default light design

### Headers Panel

- use a compact key-value list with soft separators

### Empty State

- muted illustration or icon
- clear instructional text
- optional example CTA

---

## 21. Metrics and Mini Cards

The reference image uses stat cards effectively. This project should do the same in smaller doses.

### Suggested Metric Cards

- Requests today
- Success rate
- Average response time
- Collections saved
- Remaining quota

### Style

- small rounded cards
- large numeric value
- tiny contextual caption
- optional mini icon button in top-right

Use 3 to 5 cards at most in a row. Do not over-dashboard the workspace.

---

## 22. Pattern and Chart Language

The reference image uses striped patterns and green chart shapes. Reuse that language carefully.

### Approved Pattern Usage

- striped backgrounds for pending states
- striped skeleton-like empty blocks
- soft green radial backgrounds for highlight cards

### Do Not Overuse

- striped patterns should not appear in every card
- use decorative patterns mainly in:
  - guest lock cards
  - usage cards
  - empty or pending states

---

## 23. Badges and Status Chips

Use soft pill badges throughout the product.

### Badge Variants

- Success: soft green background with dark green text
- Warning: pale amber background with darker amber text
- Danger: pale red background with darker red text
- Neutral: warm gray background with muted text

### API Method Chips

Unlike the rest of the UI, method chips may have slightly clearer distinction:

- `GET`: green tint
- `POST`: teal-green tint
- `PUT`: warm gold tint
- `PATCH`: sage tint
- `DELETE`: soft red tint

Keep them soft, not saturated.

---

## 24. Guest Lock States

Since `/app` is shared by guests and signed-in users, locked states must feel productized.

### Locked Feature Pattern

- keep the real module visible
- blur, dim, or disable only the restricted control area
- show a soft card overlay with:
  - lock icon
  - short explanation
  - sign-in or upgrade CTA

### Good Locked Areas

- save collection
- custom domain execution
- environment variables
- history persistence
- advanced snippets/tests

Do not remove these sections entirely for guests. Show the capability and gate it cleanly.

---

## 25. Motion and Interaction

Motion should be subtle and polished.

### Recommended Motion Rules

- hover transitions: `150ms` to `200ms`
- panel reveal: `200ms` to `260ms`
- button press: slight scale down or shadow reduction
- no bouncing or dramatic spring animations

### Hover Behavior

- cards: tiny lift or border emphasis
- nav items: background tint
- buttons: darker fill or border emphasis

---

## 26. Responsive Behavior

The desktop experience is the primary visual target, but mobile should remain coherent.

### Desktop

- full shell layout
- persistent sidebar
- side-by-side request and response panels

### Tablet

- sidebar can collapse
- request/response can stack or use tabs

### Mobile

- shell becomes edge-to-edge with reduced outer padding
- toolbar compresses heavily
- request and response must stack vertically
- stats row becomes horizontal scroll or 2-column grid

Do not force the exact desktop dashboard layout onto mobile.

---

## 27. Landing Page Adaptation

The landing page should use the same visual language as the app.

### Hero

- warm neutral page canvas
- large rounded product preview shell
- green CTA
- white outline secondary CTA

### Feature Sections

- use card grids similar to the app
- include mini workspace previews
- maintain consistent spacing and radius

### Key Rule

The marketing site should feel like an extension of the product, not a separate branding exercise.

---

## 28. Tailwind + Component Primitive Implementation Rules

### Rule 1
- Use theme tokens first.
- Avoid hardcoding arbitrary colors in component markup unless it is a one-off decorative case.

### Rule 2
- Customize shared component primitives to match this visual system globally.
- Button, Card, Input, Tabs, Badge, Sheet, Dialog, DropdownMenu, ScrollArea, Separator, and Tooltip should all inherit the same radius, borders, and shadow logic.

### Rule 3
- Create a small set of reusable app-level utility classes.

Recommended utilities:

```css
.app-shell
.panel-card
.soft-card
.metric-card
.section-title
.pill-button
.locked-overlay
.stripe-fill
```

### Rule 4
- Use Tailwind class composition for layout, but keep repeated visual rules in reusable classes or components.

### Rule 5
- Keep code editor surfaces visually integrated with the design system.
- Do not let Monaco or CodeMirror default theme clash with the surrounding UI.

---

## 29. Suggested Reusable Utility Classes

These are implementation-level style helpers worth creating.

```css
.app-shell {
  @apply mx-auto min-h-[calc(100vh-2rem)] max-w-[1440px] rounded-[32px] border border-white/60 bg-shell p-3 shadow-shell;
}

.panel-card {
  @apply rounded-[24px] border border-border/80 bg-card shadow-card;
}

.soft-card {
  @apply rounded-[20px] border border-border/70 bg-panel-soft;
}

.metric-card {
  @apply rounded-[20px] border border-border/80 bg-white p-5 shadow-card;
}

.pill-button {
  @apply rounded-full px-5 py-2.5 font-medium;
}

.locked-overlay {
  @apply rounded-[20px] border border-border/70 bg-white/85 backdrop-blur-sm;
}
```

---

## 30. Component-by-Component Guidance

### Button
- default variant should map to the green pill style
- outline variant should map to white bordered pill

### Card
- increase rounding beyond shadcn defaults
- use soft shadow and soft border

### Input
- increase height slightly
- round more than default
- use lighter border and placeholder

### Tabs
- use pill-like container and rounded active indicator

### Badge
- use fully rounded chip style

### Dialog / Sheet
- large radius
- white surface
- minimal shadow noise

### Tooltip / Dropdown
- soft white surface with subtle border

---

## 31. Do's

- Do keep the interface light-first and warm.
- Do use green as the main accent consistently.
- Do preserve generous rounding across the whole product.
- Do keep cards structured and modular.
- Do make technical content readable with strong spacing and monospace where needed.
- Do use locked states instead of hiding premium or auth-gated functionality.
- Do keep the landing page and app visually connected.
- Do implement the style through tokens and reusable primitives.

---

## 32. Don'ts

- Don't switch to a dark devtool aesthetic by default.
- Don't use pure white everywhere.
- Don't use sharp corners or harsh borders.
- Don't overuse gradients, blur, or glass effects.
- Don't rely on purple, blue, or neon accents as primary brand colors.
- Don't let code blocks look like they belong to a different product.
- Don't create cluttered dashboards with too many metric modules.
- Don't hardcode styling inconsistently across components.

---

## 33. Final Style Definition

The final product should feel like this:

**A rounded, warm, premium API workspace with quiet confidence, strong green accents, and technical clarity.**

If a screen looks too cold, too flat, too dark, too sharp, or too generic, it is off-style.
