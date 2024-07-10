.PHONY: build test run migration migrate-up migrate-down

build:
	@go build -o bin/didis cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/didis

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

coverage:
	@go test -coverprofile=coverage.out $(shell go list ./... | grep -v /mocks)
	@go tool cover -html=coverage.out -o coverage.html
