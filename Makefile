# Variables
DB_URL ?= postgres://gator:gator@localhost:5432/gator
MIGRATIONS_DIR = sql/schema

# Target
.PHONY: migrate

# Run all Goose UP migrations
migrate:
	goose -dir ${MIGRATIONS_DIR} postgres "${DB_URL}" up

