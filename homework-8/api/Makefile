# Параметры для подключения к базе данных PostgreSQL
PG_HOST=localhost
PG_PORT=5432
PG_USER=test
PG_PASSWORD=test
PG_DATABASE=test

ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=$(PG_USER) password=$(PG_PASSWORD) dbname=$(PG_DATABASE) host=$(PG_HOST) port=$(PG_PORT) sslmode=disable
endif

MIGRATION_FOLDER=$(CURDIR)/migrations

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: test-migration-up
test-migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: test-migration-down
test-migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down

.PHONY: generate
generate:
	rm -rf internal/pb/*.go
	
	protoc \
		--proto_path=proto/ \
		--go_out=internal/pb \
		--go-grpc_out=internal/pb \
		proto/*.proto

	
.PHONY: compose-up
compose-up:
	docker-compose build
	docker-compose up -d postgres

.PHONY: compose-rm
compose-rm:
	docker-compose down
