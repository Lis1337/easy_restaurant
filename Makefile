compose = docker compose
docker = docker

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
	$(docker) exec -it app sh
postgres-exec:
	$(docker) exec -it postgres sh
