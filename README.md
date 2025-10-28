## Task Manager API (Go + Gin + GORM + Postgres)

A simple CRUD API for managing tasks, built with Gin, GORM, and Postgres. Includes Docker Compose for the database and a clean layered structure with handlers, services, and routes.

### Tech Stack

- Go, Gin (HTTP framework)
- GORM (ORM)
- Postgres 15 (Docker)

### Project Structure

```
cmd/server/main.go         # App entrypoint
config/                    # Configuration loader (.env support)
internal/db/               # DB connection and migration (GORM)
internal/models/           # GORM models
internal/services/         # Business logic (uses DB)
internal/handlers/         # HTTP handlers (bind/validate/respond)
internal/routes/           # Route registration (Gin groups)
docker-compose.yml         # Postgres service
```

### Prerequisites

- Go 1.21+
- Docker (for Postgres)
- psql (optional, for CLI access)

### Environment Variables

The app reads variables via `config.LoadConfig()` with sane defaults:

```
DB_HOST=localhost
DB_PORT=5432
DB_NAME=taskmanager
DB_USER=postgres
DB_PASSWORD=password
```

Create a `.env` file at the repo root to override as needed. When running the Go app inside Docker, set `DB_HOST=postgres` (the Compose service name).

### Run the Database (Docker)

```bash
docker compose up -d postgres
```

Verify it’s up:

```bash
docker ps
```

Optional: connect with psql from host

```bash
psql -h localhost -p 5432 -U postgres -d taskmanager
# password: password
```

Or from inside the container:

```bash
docker exec -it taskmanager_postgres psql -U postgres -d taskmanager
```

### Run the API (Local Dev)

```bash
go mod download
go run ./cmd/server
```

Server starts on `http://localhost:8080`.

### Health Check

```bash
curl http://localhost:8080/health
```

### API Endpoints

Base path: `/api`

- POST `/api/tasks` – Create task
- GET `/api/tasks` – List tasks
- GET `/api/tasks/:id` – Get task by ID
- PUT `/api/tasks/:id` – Update task
- DELETE `/api/tasks/:id` – Delete task

#### Example Requests

Create:

```bash
curl -X POST http://localhost:8080/api/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Read docs","description":"Read the README","completed":false}'
```

List:

```bash
curl http://localhost:8080/api/tasks
```

Get by ID:

```bash
curl http://localhost:8080/api/tasks/1
```

Update:

```bash
curl -X PUT http://localhost:8080/api/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Read docs (updated)","description":"...","completed":true}'
```

Delete:

```bash
curl -X DELETE http://localhost:8080/api/tasks/1
```

### Code Overview

- `internal/routes/routes.go` registers all endpoints and groups.
- `internal/handlers/task_handlers.go` validates input and shapes responses.
- `internal/services/task_services.go` performs DB operations via GORM.
- `internal/db/connection.go` opens the DB and runs auto-migrations for `models.Task`.

### Common Issues

- Cannot connect to DB: ensure `docker compose up -d postgres` and environment variables match. For containerized app, use `DB_HOST=postgres`.
- Port in use: change host port mapping in `docker-compose.yml` or stop existing Postgres.

### License

MIT
