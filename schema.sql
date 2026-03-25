-- API Testing Kit
-- PostgreSQL schema for a guest + authenticated API testing workspace.
-- Assumes application-layer encryption for secrets and sensitive payload fragments.

CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS citext;

-- ---------------------------------------------------------------------------
-- ENUMS
-- ---------------------------------------------------------------------------

DO $$
BEGIN
  CREATE TYPE user_status AS ENUM ('active', 'invited', 'pending_verification', 'suspended', 'deleted');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
  CREATE TYPE session_status AS ENUM ('active', 'revoked', 'expired');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
  CREATE TYPE collection_visibility AS ENUM ('private', 'shared_readonly', 'internal');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
  CREATE TYPE request_method AS ENUM ('GET', 'POST', 'PUT', 'PATCH', 'DELETE', 'HEAD', 'OPTIONS');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
  CREATE TYPE request_body_mode AS ENUM ('none', 'raw', 'json', 'form_urlencoded', 'form_data');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
  CREATE TYPE auth_scheme AS ENUM ('none', 'basic', 'bearer', 'api_key', 'oauth2', 'custom');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
  CREATE TYPE run_source AS ENUM ('guest', 'authenticated', 'template', 'manual_replay', 'import');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
  CREATE TYPE run_status AS ENUM ('queued', 'running', 'succeeded', 'failed', 'blocked', 'timed_out', 'canceled');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
  CREATE TYPE template_visibility AS ENUM ('public', 'private', 'featured');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
  CREATE TYPE entitlement_source AS ENUM ('plan', 'trial', 'manual', 'comped', 'subscription');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
  CREATE TYPE subscription_status AS ENUM ('trialing', 'active', 'past_due', 'paused', 'canceled', 'incomplete', 'unpaid');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
  CREATE TYPE billing_provider AS ENUM ('stripe', 'paddle', 'lemonsqueezy', 'manual');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
  CREATE TYPE billing_interval AS ENUM ('month', 'year', 'one_time');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
  CREATE TYPE usage_bucket AS ENUM ('minute', 'hour', 'day', 'month');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
  CREATE TYPE abuse_severity AS ENUM ('low', 'medium', 'high', 'critical');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
  CREATE TYPE blocked_target_type AS ENUM ('domain', 'host', 'ip', 'cidr', 'url_pattern');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;

-- ---------------------------------------------------------------------------
-- CORE IDENTITY / AUTH
-- ---------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS users (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  email citext NOT NULL UNIQUE,
  email_verified_at timestamptz,
  password_hash text,
  display_name text,
  avatar_url text,
  status user_status NOT NULL DEFAULT 'pending_verification',
  role text NOT NULL DEFAULT 'user',
  locale text NOT NULL DEFAULT 'en',
  timezone text NOT NULL DEFAULT 'Asia/Calcutta',
  last_login_at timestamptz,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  deleted_at timestamptz
);

COMMENT ON TABLE users IS 'Primary account table for guests who upgrade and authenticated users.';
COMMENT ON COLUMN users.password_hash IS 'Store a password hash only, never a plaintext password. Consider app-layer encryption for any sensitive auth metadata.';

CREATE TABLE IF NOT EXISTS user_identities (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  provider text NOT NULL,
  provider_user_id text NOT NULL,
  provider_email text,
  access_token_encrypted text,
  refresh_token_encrypted text,
  token_expires_at timestamptz,
  profile jsonb NOT NULL DEFAULT '{}'::jsonb,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  UNIQUE (provider, provider_user_id)
);

COMMENT ON TABLE user_identities IS 'Optional OAuth or external identity links. Token fields should be encrypted in the application layer.';

CREATE TABLE IF NOT EXISTS sessions (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  session_token_hash text NOT NULL UNIQUE,
  status session_status NOT NULL DEFAULT 'active',
  ip_address inet,
  user_agent text,
  expires_at timestamptz NOT NULL,
  last_seen_at timestamptz,
  revoked_at timestamptz,
  created_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_sessions_user_id_status ON sessions (user_id, status);
CREATE INDEX IF NOT EXISTS idx_sessions_expires_at ON sessions (expires_at);

CREATE TABLE IF NOT EXISTS email_verification_tokens (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  token_hash text NOT NULL UNIQUE,
  email text NOT NULL,
  expires_at timestamptz NOT NULL,
  consumed_at timestamptz,
  created_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_email_verification_tokens_user_id ON email_verification_tokens (user_id);
CREATE INDEX IF NOT EXISTS idx_email_verification_tokens_expires_at ON email_verification_tokens (expires_at);

CREATE TABLE IF NOT EXISTS password_reset_tokens (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  token_hash text NOT NULL UNIQUE,
  expires_at timestamptz NOT NULL,
  consumed_at timestamptz,
  created_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE password_reset_tokens IS 'Optional local-auth support if password login is enabled.';

-- ---------------------------------------------------------------------------
-- WORKSPACES / REQUEST BUILDING
-- ---------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS collections (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  owner_user_id uuid REFERENCES users(id) ON DELETE SET NULL,
  name text NOT NULL,
  slug text,
  description text,
  visibility collection_visibility NOT NULL DEFAULT 'private',
  color text,
  sort_order integer NOT NULL DEFAULT 0,
  shared_token text UNIQUE,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  deleted_at timestamptz,
  CONSTRAINT collections_name_not_empty CHECK (length(trim(name)) > 0)
);

COMMENT ON TABLE collections IS 'Saved request groups shown in the app sidebar and collection pages.';

CREATE INDEX IF NOT EXISTS idx_collections_owner_user_id ON collections (owner_user_id);
CREATE INDEX IF NOT EXISTS idx_collections_visibility ON collections (visibility);

CREATE TABLE IF NOT EXISTS saved_requests (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  collection_id uuid REFERENCES collections(id) ON DELETE CASCADE,
  owner_user_id uuid REFERENCES users(id) ON DELETE SET NULL,
  name text NOT NULL,
  description text,
  method request_method NOT NULL DEFAULT 'GET',
  url text NOT NULL,
  path_template text,
  query_params jsonb NOT NULL DEFAULT '[]'::jsonb,
  headers jsonb NOT NULL DEFAULT '[]'::jsonb,
  auth_scheme auth_scheme NOT NULL DEFAULT 'none',
  auth_config jsonb NOT NULL DEFAULT '{}'::jsonb,
  body_mode request_body_mode NOT NULL DEFAULT 'none',
  body_config jsonb NOT NULL DEFAULT '{}'::jsonb,
  pre_request_script text,
  post_request_script text,
  assertions jsonb NOT NULL DEFAULT '[]'::jsonb,
  example_response jsonb NOT NULL DEFAULT '{}'::jsonb,
  tags text[] NOT NULL DEFAULT '{}'::text[],
  is_template boolean NOT NULL DEFAULT false,
  sort_order integer NOT NULL DEFAULT 0,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  deleted_at timestamptz,
  CONSTRAINT saved_requests_name_not_empty CHECK (length(trim(name)) > 0),
  CONSTRAINT saved_requests_url_http CHECK (url ~* '^https?://')
);

COMMENT ON TABLE saved_requests IS 'Reusable API request definitions. Sensitive auth values should be stored encrypted or referenced indirectly.';
COMMENT ON COLUMN saved_requests.auth_config IS 'May contain tokens, secrets, or references. Use application-layer encryption for secret values.';
COMMENT ON COLUMN saved_requests.body_config IS 'Stores JSON/raw/form-data configuration and any structured editor state.';

CREATE INDEX IF NOT EXISTS idx_saved_requests_collection_id ON saved_requests (collection_id);
CREATE INDEX IF NOT EXISTS idx_saved_requests_owner_user_id ON saved_requests (owner_user_id);
CREATE INDEX IF NOT EXISTS idx_saved_requests_is_template ON saved_requests (is_template);

CREATE TABLE IF NOT EXISTS template_collections (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  slug text NOT NULL UNIQUE,
  name text NOT NULL,
  description text,
  visibility template_visibility NOT NULL DEFAULT 'public',
  category text,
  featured_order integer NOT NULL DEFAULT 0,
  is_active boolean NOT NULL DEFAULT true,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE template_collections IS 'Curated guest-safe template groups shown on the templates page and inside /app.';

CREATE INDEX IF NOT EXISTS idx_template_collections_visibility ON template_collections (visibility);
CREATE INDEX IF NOT EXISTS idx_template_collections_is_active ON template_collections (is_active);

CREATE TABLE IF NOT EXISTS request_templates (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  template_collection_id uuid REFERENCES template_collections(id) ON DELETE SET NULL,
  slug text NOT NULL UNIQUE,
  name text NOT NULL,
  description text,
  visibility template_visibility NOT NULL DEFAULT 'public',
  category text,
  request_definition jsonb NOT NULL,
  preview_response jsonb NOT NULL DEFAULT '{}'::jsonb,
  sort_order integer NOT NULL DEFAULT 0,
  featured_order integer NOT NULL DEFAULT 0,
  is_active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE request_templates IS 'Curated guest-safe examples displayed in the public app experience.';

CREATE INDEX IF NOT EXISTS idx_request_templates_visibility ON request_templates (visibility);
CREATE INDEX IF NOT EXISTS idx_request_templates_is_active ON request_templates (is_active);
CREATE INDEX IF NOT EXISTS idx_request_templates_template_collection_id ON request_templates (template_collection_id);

CREATE TABLE IF NOT EXISTS request_runs (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid REFERENCES users(id) ON DELETE SET NULL,
  collection_id uuid REFERENCES collections(id) ON DELETE SET NULL,
  saved_request_id uuid REFERENCES saved_requests(id) ON DELETE SET NULL,
  template_id uuid REFERENCES request_templates(id) ON DELETE SET NULL,
  source run_source NOT NULL DEFAULT 'authenticated',
  status run_status NOT NULL DEFAULT 'queued',
  method request_method NOT NULL,
  url text NOT NULL,
  final_url text,
  target_host text,
  target_ip inet,
  request_headers jsonb NOT NULL DEFAULT '[]'::jsonb,
  request_query_params jsonb NOT NULL DEFAULT '[]'::jsonb,
  request_auth jsonb NOT NULL DEFAULT '{}'::jsonb,
  request_body jsonb NOT NULL DEFAULT '{}'::jsonb,
  request_body_preview text,
  request_size_bytes integer,
  response_status integer,
  response_headers jsonb NOT NULL DEFAULT '[]'::jsonb,
  response_body text,
  response_body_json jsonb,
  response_body_preview text,
  response_size_bytes integer,
  response_time_ms integer,
  response_content_type text,
  redirect_count integer NOT NULL DEFAULT 0,
  blocked_reason text,
  error_code text,
  error_message text,
  started_at timestamptz,
  completed_at timestamptz,
  created_at timestamptz NOT NULL DEFAULT now(),
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  CONSTRAINT request_runs_url_http CHECK (url ~* '^https?://'),
  CONSTRAINT request_runs_final_url_http CHECK (final_url IS NULL OR final_url ~* '^https?://')
);

COMMENT ON TABLE request_runs IS 'Executed request history and response snapshots. Store only what you need for replay/debugging.';
COMMENT ON COLUMN request_runs.request_auth IS 'May reference encrypted secrets or opaque credential identifiers.';
COMMENT ON COLUMN request_runs.response_body IS 'Consider response truncation limits and retention policies.';

CREATE INDEX IF NOT EXISTS idx_request_runs_user_id_created_at ON request_runs (user_id, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_request_runs_collection_id_created_at ON request_runs (collection_id, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_request_runs_saved_request_id_created_at ON request_runs (saved_request_id, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_request_runs_status_created_at ON request_runs (status, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_request_runs_target_host ON request_runs (target_host);

CREATE TABLE IF NOT EXISTS environments (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  owner_user_id uuid REFERENCES users(id) ON DELETE CASCADE,
  collection_id uuid REFERENCES collections(id) ON DELETE CASCADE,
  name text NOT NULL,
  description text,
  is_default boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  deleted_at timestamptz,
  CONSTRAINT environments_one_owner_or_collection CHECK (
    (owner_user_id IS NOT NULL AND collection_id IS NULL)
    OR
    (owner_user_id IS NULL AND collection_id IS NOT NULL)
  )
);

CREATE INDEX IF NOT EXISTS idx_environments_owner_user_id ON environments (owner_user_id);
CREATE INDEX IF NOT EXISTS idx_environments_collection_id ON environments (collection_id);

CREATE TABLE IF NOT EXISTS environment_variables (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  environment_id uuid NOT NULL REFERENCES environments(id) ON DELETE CASCADE,
  name text NOT NULL,
  value_encrypted text,
  value_preview text,
  is_secret boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  UNIQUE (environment_id, name)
);

COMMENT ON TABLE environment_variables IS 'Store secrets encrypted in the application layer. Keep previews non-sensitive.';
COMMENT ON COLUMN environment_variables.value_encrypted IS 'Application-layer encrypted secret value.';

CREATE INDEX IF NOT EXISTS idx_environment_variables_environment_id ON environment_variables (environment_id);

-- ---------------------------------------------------------------------------
-- USAGE / ABUSE / SAFETY
-- ---------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS usage_events (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid REFERENCES users(id) ON DELETE SET NULL,
  session_id uuid REFERENCES sessions(id) ON DELETE SET NULL,
  request_run_id uuid REFERENCES request_runs(id) ON DELETE SET NULL,
  bucket usage_bucket NOT NULL DEFAULT 'month',
  event_key text NOT NULL,
  quantity integer NOT NULL DEFAULT 1,
  dimensions jsonb NOT NULL DEFAULT '{}'::jsonb,
  occurred_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE usage_events IS 'Time-series usage records for quotas, analytics, and plan enforcement.';

CREATE INDEX IF NOT EXISTS idx_usage_events_user_id_occurred_at ON usage_events (user_id, occurred_at DESC);
CREATE INDEX IF NOT EXISTS idx_usage_events_event_key_occurred_at ON usage_events (event_key, occurred_at DESC);
CREATE INDEX IF NOT EXISTS idx_usage_events_occurred_at ON usage_events (occurred_at DESC);

CREATE TABLE IF NOT EXISTS abuse_events (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid REFERENCES users(id) ON DELETE SET NULL,
  session_id uuid REFERENCES sessions(id) ON DELETE SET NULL,
  request_run_id uuid REFERENCES request_runs(id) ON DELETE SET NULL,
  severity abuse_severity NOT NULL DEFAULT 'low',
  category text NOT NULL,
  source_ip inet,
  target text,
  rule_key text,
  action_taken text NOT NULL,
  details jsonb NOT NULL DEFAULT '{}'::jsonb,
  created_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE abuse_events IS 'Audit trail for blocked requests, suspicious behavior, and automated enforcement.';

CREATE INDEX IF NOT EXISTS idx_abuse_events_user_id_created_at ON abuse_events (user_id, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_abuse_events_category_created_at ON abuse_events (category, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_abuse_events_source_ip_created_at ON abuse_events (source_ip, created_at DESC);

CREATE TABLE IF NOT EXISTS blocked_targets (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  target_type blocked_target_type NOT NULL,
  target_value text NOT NULL,
  reason text NOT NULL,
  source text NOT NULL DEFAULT 'manual',
  is_active boolean NOT NULL DEFAULT true,
  expires_at timestamptz,
  created_by_user_id uuid REFERENCES users(id) ON DELETE SET NULL,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  UNIQUE (target_type, target_value)
);

COMMENT ON TABLE blocked_targets IS 'Domain/IP/path denylist used to prevent SSRF and abuse.';

CREATE INDEX IF NOT EXISTS idx_blocked_targets_is_active ON blocked_targets (is_active);
CREATE INDEX IF NOT EXISTS idx_blocked_targets_expires_at ON blocked_targets (expires_at);

-- ---------------------------------------------------------------------------
-- BILLING / ENTITLEMENTS
-- ---------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS plans (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  code text NOT NULL UNIQUE,
  name text NOT NULL,
  description text,
  is_active boolean NOT NULL DEFAULT true,
  sort_order integer NOT NULL DEFAULT 0,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE plans IS 'Canonical product plan definitions independent of the billing provider.';

CREATE TABLE IF NOT EXISTS plan_prices (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  plan_id uuid NOT NULL REFERENCES plans(id) ON DELETE CASCADE,
  provider billing_provider NOT NULL,
  provider_price_id text NOT NULL,
  billing_interval billing_interval NOT NULL DEFAULT 'month',
  amount_cents integer NOT NULL DEFAULT 0 CHECK (amount_cents >= 0),
  currency text NOT NULL DEFAULT 'usd',
  is_active boolean NOT NULL DEFAULT true,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  UNIQUE (provider, provider_price_id)
);

COMMENT ON TABLE plan_prices IS 'Provider-specific sellable price points for a plan. Supports one provider now and leaves room for another later.';

CREATE INDEX IF NOT EXISTS idx_plan_prices_plan_id ON plan_prices (plan_id);
CREATE INDEX IF NOT EXISTS idx_plan_prices_provider_is_active ON plan_prices (provider, is_active);

CREATE TABLE IF NOT EXISTS plan_entitlements (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  plan_id uuid NOT NULL REFERENCES plans(id) ON DELETE CASCADE,
  key text NOT NULL,
  value jsonb NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  UNIQUE (plan_id, key)
);

COMMENT ON TABLE plan_entitlements IS 'Plan-level feature limits and capabilities such as quotas, request execution, and retention.';

CREATE INDEX IF NOT EXISTS idx_plan_entitlements_plan_id ON plan_entitlements (plan_id);

CREATE TABLE IF NOT EXISTS billing_customers (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  provider billing_provider NOT NULL,
  provider_customer_id text NOT NULL,
  provider_email text,
  tax_country text,
  tax_id_last4 text,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  UNIQUE (provider, provider_customer_id),
  UNIQUE (user_id, provider)
);

COMMENT ON TABLE billing_customers IS 'Maps internal users to billing provider customers. Store only non-sensitive provider metadata here.';

CREATE TABLE IF NOT EXISTS subscriptions (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  billing_customer_id uuid REFERENCES billing_customers(id) ON DELETE SET NULL,
  plan_id uuid NOT NULL REFERENCES plans(id) ON DELETE RESTRICT,
  plan_price_id uuid REFERENCES plan_prices(id) ON DELETE SET NULL,
  provider billing_provider NOT NULL,
  provider_subscription_id text NOT NULL,
  status subscription_status NOT NULL DEFAULT 'incomplete',
  quantity integer NOT NULL DEFAULT 1,
  trial_ends_at timestamptz,
  current_period_start timestamptz,
  current_period_end timestamptz,
  cancel_at_period_end boolean NOT NULL DEFAULT false,
  canceled_at timestamptz,
  paused_at timestamptz,
  past_due_since timestamptz,
  payment_method_hint text,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  UNIQUE (provider, provider_subscription_id)
);

COMMENT ON TABLE subscriptions IS 'Provider-backed subscription state. Webhooks should be the source of truth.';

CREATE INDEX IF NOT EXISTS idx_subscriptions_user_id_status ON subscriptions (user_id, status);
CREATE INDEX IF NOT EXISTS idx_subscriptions_billing_customer_id ON subscriptions (billing_customer_id);

CREATE TABLE IF NOT EXISTS subscription_events (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  subscription_id uuid NOT NULL REFERENCES subscriptions(id) ON DELETE CASCADE,
  provider billing_provider NOT NULL,
  provider_event_id text NOT NULL,
  event_type text NOT NULL,
  payload jsonb NOT NULL,
  processed_at timestamptz,
  created_at timestamptz NOT NULL DEFAULT now(),
  UNIQUE (provider, provider_event_id)
);

COMMENT ON TABLE subscription_events IS 'Webhook/event inbox for Stripe, Paddle, or Lemon integrations.';

CREATE INDEX IF NOT EXISTS idx_subscription_events_subscription_id_created_at ON subscription_events (subscription_id, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_subscription_events_event_type_created_at ON subscription_events (event_type, created_at DESC);

CREATE TABLE IF NOT EXISTS invoices (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  subscription_id uuid REFERENCES subscriptions(id) ON DELETE SET NULL,
  provider billing_provider NOT NULL,
  provider_invoice_id text NOT NULL,
  invoice_number text,
  status text NOT NULL,
  amount_due_cents integer NOT NULL DEFAULT 0 CHECK (amount_due_cents >= 0),
  amount_paid_cents integer NOT NULL DEFAULT 0 CHECK (amount_paid_cents >= 0),
  currency text NOT NULL DEFAULT 'usd',
  hosted_invoice_url text,
  pdf_url text,
  issued_at timestamptz,
  due_at timestamptz,
  paid_at timestamptz,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  UNIQUE (provider, provider_invoice_id)
);

COMMENT ON TABLE invoices IS 'Optional invoice history for billing portal and support workflows.';

CREATE INDEX IF NOT EXISTS idx_invoices_subscription_id ON invoices (subscription_id);
CREATE INDEX IF NOT EXISTS idx_invoices_status_issued_at ON invoices (status, issued_at DESC);

CREATE TABLE IF NOT EXISTS user_entitlements (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  plan_id uuid REFERENCES plans(id) ON DELETE SET NULL,
  source entitlement_source NOT NULL DEFAULT 'manual',
  key text NOT NULL,
  value jsonb NOT NULL,
  starts_at timestamptz,
  ends_at timestamptz,
  is_active boolean NOT NULL DEFAULT true,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  UNIQUE (user_id, key, source)
);

COMMENT ON TABLE user_entitlements IS 'Flattened effective entitlements for gating features and quotas at runtime.';

CREATE INDEX IF NOT EXISTS idx_user_entitlements_user_id_is_active ON user_entitlements (user_id, is_active);
CREATE INDEX IF NOT EXISTS idx_user_entitlements_key_is_active ON user_entitlements (key, is_active);

-- ---------------------------------------------------------------------------
-- OPTIONAL AUDIT / OPERATIONS
-- ---------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS audit_log_entries (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  actor_user_id uuid REFERENCES users(id) ON DELETE SET NULL,
  entity_type text NOT NULL,
  entity_id uuid,
  action text NOT NULL,
  before_state jsonb NOT NULL DEFAULT '{}'::jsonb,
  after_state jsonb NOT NULL DEFAULT '{}'::jsonb,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  created_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE audit_log_entries IS 'Generic admin audit trail for account, billing, and abuse operations.';

CREATE INDEX IF NOT EXISTS idx_audit_log_entries_actor_user_id_created_at ON audit_log_entries (actor_user_id, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_audit_log_entries_entity_type_entity_id ON audit_log_entries (entity_type, entity_id);

-- ---------------------------------------------------------------------------
-- UPDATED_AT TRIGGERS
-- ---------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS trigger
LANGUAGE plpgsql
AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$;

DO $$
DECLARE
  t text;
BEGIN
  FOREACH t IN ARRAY ARRAY[
    'users',
    'user_identities',
    'collections',
    'saved_requests',
    'template_collections',
    'request_templates',
    'environments',
    'environment_variables',
    'blocked_targets',
    'plans',
    'plan_prices',
    'subscriptions',
    'invoices',
    'user_entitlements'
  ]
  LOOP
    EXECUTE format('DROP TRIGGER IF EXISTS trg_%I_updated_at ON %I;', t, t);
    EXECUTE format('CREATE TRIGGER trg_%I_updated_at BEFORE UPDATE ON %I FOR EACH ROW EXECUTE FUNCTION set_updated_at();', t, t);
  END LOOP;
END $$;
