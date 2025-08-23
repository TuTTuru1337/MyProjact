.PHONY: migrate migrate-new migrate-down

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
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

lint:
	golangci-lint run --color=auto

test:
	go test ./... -v

run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server