GOTOOLS := \
	github.com/alecthomas/gometalinter \
	github.com/golang/dep/cmd/dep

DATABASE_NAME := $(shell cat pkg/config/development.json | jq '.["database_name"]')
DATABASE_USER := $(shell cat pkg/config/development.json | jq '.["database_user"]') 

.PHONY: tools
tools: ## First-time installation of development tools
	go get -u $(GOTOOLS)
	gometalinter --install

.PHONY: deps
deps: ## Ensure all dependencies are in sync
	dep ensure

.PHONY: db-setup
db-setup: ## Create database
	createuser --createdb $(DATABASE_USER)
	createdb -U $(DATABASE_USER) $(DATABASE_NAME)
	make db-migrate

.PHONY: db-migrate
db-migrate: ## Run new database migrations
	go run cmd/migrations/*.go up

.PHONY: db-rollback
db-rollback: ## Rollback last run database migrations
	go run cmd/migrations/*.go down

.PHONY: db-migrate-create
db-migrate-create: ## Add a migration: make db-migrate-create name=add_table
	go run cmd/migrations/*.go create $(name)

.PHONY: db-reset
db-reset: ## Delete and recreate database
	dropdb --if-exists $(DATABASE_NAME)
	dropuser --if-exists $(DATABASE_USER)
	make db-setup

.PHONY: run
run: ## Run application
	gin --build ./cmd/serve --immediate run

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := run
