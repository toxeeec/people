DOCKER_COMPOSE_DEV=compose.dev.yaml
DOCKER_COMPOSE_TEST=compose.test.yaml
PROJECT_TEST=people-test
BACKEND_SERVICE=backend

all: up

up:
	docker compose -f $(DOCKER_COMPOSE_DEV) up

down:
	docker compose -f $(DOCKER_COMPOSE_DEV) down
	docker compose -f $(DOCKER_COMPOSE_TEST) down

test:
	docker compose -f $(DOCKER_COMPOSE_TEST) -p $(PROJECT_TEST) build
	docker compose -f $(DOCKER_COMPOSE_TEST) -p $(PROJECT_TEST) run $(BACKEND_SERVICE); docker compose -f $(DOCKER_COMPOSE_TEST) -p $(PROJECT_TEST) down
