install:
	go mod tidy
	go install github.com/rubenv/sql-migrate/...@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	cd config && sqlc generate

swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g ./cmd/main.go

generate:
	cd config && sqlc generate

migrate:
	sql-migrate up -config ./config/sql-migrate.yml -env default

unmigrate-1:
	sql-migrate down -config ./config/sql-migrate.yml -env default -limit 1

seed:
	sql-migrate up -config ./config/sql-seed.yml -env default

run:
	go run ./cmd/main.go

build:
	sh ./scripts/build.sh
