# Makefile for database migrations with golang-migrate

# Load environment variables from .env
include .env
export $(shell sed 's/=.*//' .env)

MIGRATE=migrate -path migrations -database "$(POSTGRES_URI)"

.PHONY: migrate-up migrate-down migrate-force migrate-create

## Apply all up migrations
migrate-up:
	$(MIGRATE) up

## Rollback the last migration
migrate-down:
	$(MIGRATE) down 1

## Force migration version (useful for fixing failed migrations)
migrate-force:
	$(MIGRATE) force $(VERSION)

## Create a new migration file
migrate-create:
	migrate create -ext sql -dir migrations -seq $(NAME)

## Show current migration status
migrate-status:
	$(MIGRATE) version
