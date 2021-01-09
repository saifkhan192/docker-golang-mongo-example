.PHONY: _list
_list:
	@echo "Type make then a space then hit tab to see available commands"

build:
	cp -u .env.dist .env
	docker-compose up -d --build
	@sleep 1

run_app:
	@google-chrome http://localhost:3000/users
	@docker logs --follow golang-app

run:
	make build
	make run_app

down:
	docker-compose down

bash_app:
	docker exec -it golang-app /bin/bash

recreate_app:
	docker-compose up -d --force-recreate --no-deps --build backend-service
	docker logs --follow golang-app

application:
	docker-compose up


run_all:
	docker container start postgresdb || true
	docker container start mongodb || true
	docker container start golang-app || true
	@google-chrome http://localhost:3000/users

stop_all:
	docker container stop postgresdb || true
	docker container stop mongodb || true
	docker container stop golang-app || true
