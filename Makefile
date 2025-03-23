compose = docker compose
docker = docker
debug = docker compose -f compose-debug.yaml

up:
	$(compose) up -d
down:
	$(compose) down
build:
	$(compose) build
dry-build:
	$(compose) build --dry-run
rebuild: down build up
recreate: down up

app-exec:
	$(docker) exec -it app bash
postgres-exec:
	$(docker) exec -it postgres sh

debug-up:
	$(debug) up -d
debug-build:
	$(debug) build
debug-down:
	$(debug) down
debug-recreate: debug-down debug-up
debug-rebuild: debug-down debug-build debug-up
