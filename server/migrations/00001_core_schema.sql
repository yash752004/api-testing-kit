-- +goose Up
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS citext;

-- +goose StatementBegin
DO $$
BEGIN
  CREATE TYPE user_status AS ENUM ('active', 'invited', 'pending_verification', 'suspended', 'deleted');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;
-- +goose StatementEnd

-- +goose StatementBegin
DO $$
BEGIN
  CREATE TYPE session_status AS ENUM ('active', 'revoked', 'expired');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;
-- +goose StatementEnd

-- +goose StatementBegin
DO $$
BEGIN
  CREATE TYPE template_visibility AS ENUM ('public', 'private', 'featured');
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;
-- +goose StatementEnd

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

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS trigger
LANGUAGE plpgsql
AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$;
-- +goose StatementEnd

DROP TRIGGER IF EXISTS trg_users_updated_at ON users;
CREATE TRIGGER trg_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

DROP TRIGGER IF EXISTS trg_template_collections_updated_at ON template_collections;
CREATE TRIGGER trg_template_collections_updated_at
BEFORE UPDATE ON template_collections
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

DROP TRIGGER IF EXISTS trg_request_templates_updated_at ON request_templates;
CREATE TRIGGER trg_request_templates_updated_at
BEFORE UPDATE ON request_templates
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

-- +goose Down
DROP TRIGGER IF EXISTS trg_request_templates_updated_at ON request_templates;
DROP TRIGGER IF EXISTS trg_template_collections_updated_at ON template_collections;
DROP TRIGGER IF EXISTS trg_users_updated_at ON users;

DROP FUNCTION IF EXISTS set_updated_at();

DROP TABLE IF EXISTS request_templates;
DROP TABLE IF EXISTS template_collections;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS template_visibility;
DROP TYPE IF EXISTS session_status;
DROP TYPE IF EXISTS user_status;

DROP EXTENSION IF EXISTS citext;
DROP EXTENSION IF EXISTS pgcrypto;
