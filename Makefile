APP_NAME = ./cmd/server/main.go
GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING ?= root:root1234@tcp(127.0.0.1:3306)/user_db
GOOSE_MIGRATION_DIR ?= sql/schema

dev:
	go run ${APP_NAME}

docker_up:
	docker compose -f docker-compose.yml up -d
docker_down:
	docker compose -f docker-compose.yml down --remove-orphans


upse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) up
downse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) down
resetse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) reset
status:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) status

.PHONY: dev docker_up docker_down upse downse resetse status