# Generate proto file

    protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/user.proto

# Run client

    go build cmd/client/main.go

# Environment variables

The server requires the following environment variables to connect to PostgreSQL:

| Variable      | Description              | Example    |
|---------------|--------------------------|------------|
| `DB_USER`     | PostgreSQL username      | `postgres` |
| `DB_PASSWORD` | PostgreSQL password      | (secret)   |

Copy `.env.example` to `.env` and fill in real values before starting the server.

# Run server

    export $(cat .env | xargs)
    go build cmd/server/main.go
