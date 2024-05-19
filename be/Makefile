install:
	go mod tidy
	go install github.com/rubenv/sql-migrate/...@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	cd config && sqlc generate

generate:
	cd config && sqlc generate

migrate:
	sql-migrate up -config ./config/sql-migrate.yml -env local

unmigrate-1:
	sql-migrate down -config ./config/sql-migrate.yml -env server -limit 1

run:
	go run ./cmd/main.go

build:
	sh ./scripts/build.sh
