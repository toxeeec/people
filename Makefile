DOCKER_COMPOSE_DEV=compose.dev.yaml

all: dev

dev:
	docker compose -f $(DOCKER_COMPOSE_DEV) up
