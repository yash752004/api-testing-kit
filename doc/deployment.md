# Deployment Guide

This repository is designed to run in a small, ordered stack:

- `db` owns PostgreSQL persistence only.
- `migrate` applies the goose migration set before application traffic starts.
- `api` serves the Go HTTP API, auth/session flow, template data, and outbound request execution.
- `web` serves the SvelteKit UI and performs server-side fetches against the API.

## Environment

Set these values for local or deployed environments:

- `DATABASE_URL` for the API and migration job.
- `API_PORT` for the Go API listener.
- `WEB_PORT` for the SvelteKit preview listener.
- `INTERNAL_API_BASE_URL` for server-side fetches from `web` to `api`.
- `PUBLIC_API_BASE_URL` for browser-visible API links and client-side navigation.

The defaults in `.env.example` are tuned for local Docker development. In production, replace the public URL with the real API origin or the private service address used by your platform.

## Local Compose Flow

Use this order when bringing the stack up from a clean checkout:

1. `docker compose up -d db`
2. `docker compose --profile migrations run --rm migrate`
3. `docker compose up -d api`
4. `docker compose up -d web`

The compose file includes health checks so `api` waits for PostgreSQL and `web` waits for the API before starting.

## Production Build

For a production-like Docker deployment, build and run the same services with the same order:

1. Build the images with `docker compose build`.
2. Run migrations once against the target database.
3. Start `db`, then `api`, then `web`.
4. Verify `GET /healthz` on the API and the web root before exposing traffic.

If you split the stack across services or platforms, keep the same boundaries:

- PostgreSQL stays isolated behind `DATABASE_URL`.
- The API remains the only component that talks to PostgreSQL directly.
- The web app only reaches the API through `INTERNAL_API_BASE_URL` on the server side.
