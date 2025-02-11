POSTGRES_PATH:="db/migrations/postgres/"
POSTGRES_DSN:="postgres://postgres:postgres@localhost:5432/ush_db?sslmode=disable" #EXAMPLE

PG_DBNAME="ush_db"

# Цель по умолчанию
.PHONY: all
all: help

.PHONY: help
help:
	@echo "Available commands:"
	@echo "migrate_pg  	-- Run migrations for Postgres"
	@echo "reset_pg		-- Delete db, create again, and do migrate_pg"

.PHONY: migrate_pg
migrate_pg:
	@echo "Run postgres migrations..."
	goose -dir $(POSTGRES_PATH) postgres $(POSTGRES_DSN) up
	@echo "Complete postgres migration"

.PHONY: reset_pg
reset_pg:
	@echo "Drop postgres db"
	goose -dir $(POSTGRES_PATH) postgres $(POSTGRES_DSN) reset

	@echo "Create postgres table"
	psql -c "create database $(PG_DBNAME)"

	@echo "Run postgres migrations..."
	goose -dir $(POSTGRES_PATH) postgres $(POSTGRES_DSN) up
	@echo "Complete postgres migration"