.PHONY: build test run migration migrate-up migrate-down

build:
	@go build -o bin/didis cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/didis

# migration:
# 	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

# migrate-up:
# 	@go run cmd/migrate/main.go up

# migrate-down:
# 	@go run cmd/migrate/main.go down

coverage:
	@go test -coverprofile=coverage.out $(shell go list ./... | grep -v /mocks)
	@go tool cover -html=coverage.out -o coverage.html


# Define variables
MIGRATIONS_DIR := cmd/migrate/migrations
NODE := node
MIGRATION_SCRIPTS := $(wildcard $(MIGRATIONS_DIR)/*.js)

# Crear una nueva migración (script de ejemplo)
create-migration:
	@echo "Creating new migration..."
	@read -p "Enter migration name: " name; \
	touch $(MIGRATIONS_DIR)/$${name}.js; \
	echo "Migration $${name}.js created."

# Ejecutar migraciones
migrate-up:
	@for file in $(MIGRATION_SCRIPTS); do \
		echo "Running $$file..."; \
		$(NODE) $$file; \
	done

# Deshacer migraciones (para simplificar, solo proporciona un comando de ejemplo)
migrate-down:
	@echo "Rolling back migrations is not supported. Manual rollback is required."