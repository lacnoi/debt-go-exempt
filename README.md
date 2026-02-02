# debt-go-exempt

go run ./cmd/server/main.go

docker create --name postgres-db -p 5432:5432 -e POSTGRES_USER=appuser -e POSTGRES_PASSWORD=apppassword -e POSTGRES_DB=appdb -v pgdata:/var/lib/postgresql/data  postgres:16

docker run -d --name postgres-db -p 5432:5432 -e POSTGRES_USER=appuser -e POSTGRES_PASSWORD=apppassword -e POSTGRES_DB=appdb -v pgdata:/var/lib/postgresql/data postgres:16