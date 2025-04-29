# Makefile

# Variables
MIGRATIONS_DIR=./migrations
BIN_NAME=todo
MAIN_PATH=./main.go

.PHONY: help build migrate_up_all install_migrate

help: ## Display a list of available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build:
	@go build -o ./bin/$(BIN_NAME) $(MAIN_PATH)

migrate_up_all: install_migrate ## Execute all `up` migrations using `migrate`
	@migrate -source file://$(MIGRATIONS_DIR) -database sqlite3://$(DB_DATASOURCE_NAME) -verbose up

install_migrate: ## Install `migrate` CLI tool
	@command -v migrate >/dev/null 2>&1 || go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
