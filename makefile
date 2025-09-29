.PHONY: migrate migrate-new migrate-down gen gen-users gen-all lint test run build generate

DB_DSN := postgres://postgres:pass1234@localhost:5432/postgres?sslmode=disable
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	@NAME=$$NAME; \
	if [ -z "$$NAME" ]; then echo "Please pass NAME=your_migration_name"; exit 1; fi; \
	migrate create -ext sql -dir ./migrations $$NAME

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

gen:
	oapi-codegen -generate types,server -package tasks ./openapi/tasks-only.yaml > ./internal/web/tasks/api.gen.go

gen-users:
	oapi-codegen -generate types,server -package users ./openapi/users-only.yaml > ./internal/web/users/api.gen.go

gen-all: gen gen-users

generate: gen-all

lint: generate
	golangci-lint run --color=auto

test:
	go test ./... -v

run: generate
	go run ./cmd/server

build: generate
	go build -o bin/server ./cmd/server