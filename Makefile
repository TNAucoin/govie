# targets not related to files
.PHONY: build up down rebuild

# environment variable
COMPOSE_FILE := ./docker-compose.yml

# make build target
build:
	docker-compose -f $(COMPOSE_FILE) build

# make up target
up:
	docker-compose -f $(COMPOSE_FILE) up

# make down target
down:
	docker-compose -f $(COMPOSE_FILE) down

rebuild:
	docker compose -f $(COMPOSE_FILE) build && docker compose -f $(COMPOSE_FILE) up