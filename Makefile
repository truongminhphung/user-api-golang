APP_NAME = ./cmd/server/main.go

dev:
	go run ${APP_NAME}

docker_up:
	docker compose -f docker-compose.yml up -d