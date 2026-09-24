# Fiber Template

Simple [Fiber](https://gofiber.io) starter with HTML templates, includes, Swagger, GORM, and SQLite.

## Run

```bash
go run ./cmd/app
```

- App: http://localhost:3000
- API: http://localhost:3000/api/v1/users
- Swagger: http://localhost:3000/swagger/

Default login: `admin@example.com` / `admin`

SQLite is stored in `app.db`. Session cookie is `session_id`.

## Layout

```text
cmd/app/                     # process entrypoint
internal/app/                # Fiber bootstrap and module wiring
internal/database/           # GORM + SQLite
internal/modules/auth/       # login, session, accounts
internal/modules/user/       # users HTML + API
internal/shared/             # shared HTTP types
internal/web/                # 404
views/auth/                  # auth templates
views/user/                  # user templates
views/layouts/               # page shell, uses {{embed}}
views/partials/              # included with {{template "partials/..." .}}
docs/                        # generated Swagger spec
```

Regenerate API docs after handler comment changes:

```bash
swag init -g main.go -d cmd/app,internal/modules/auth,internal/modules/user,internal/shared -o docs
```
