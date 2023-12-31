PROJECT_NAME = go31ngo

CONTAINER = $$(docker ps | grep ${PROJECT_NAME} | awk '{print $$1}')
DOCKER_COMPOSE := --env-file .env -p ${PROJECT_NAME} -f ops/docker/docker-compose.yml

## Builds the container image
build: copy-env delete-mongodb-data compose
	
copy-env:
	sh ops/scripts/copy-env.sh
	
compose:
	docker-compose ${DOCKER_COMPOSE} up -d --build

delete-mongodb-data:
	sh ops/scripts/delete-mongodb-data.sh

## Starts the container
start:
	docker-compose ${DOCKER_COMPOSE} up -d

restart: stop start

## Create app file
create-app:
	sh ops/scripts/create-app.sh

## Stops the container
stop: is-running
	docker-compose ${DOCKER_COMPOSE} stop

## Attach shell to the container that is running
enter: is-running
	@docker exec -it ${CONTAINER} /bin/bash

## Check if the container is running
is-running:
	@docker exec ${CONTAINER} true 2>/dev/null || (echo "Docker container is not running - Please execute ---> make start or make build <--- to start it"; exit 1)

# COLORS
TPUT := $(shell command -v tput 2> /dev/null)

ifdef TPUT
	GREEN  := $(shell tput -Txterm setaf 2)
	YELLOW := $(shell tput -Txterm setaf 3)
	WHITE  := $(shell tput -Txterm setaf 7)
	RESET  := $(shell tput -Txterm sgr0)
endif

TARGET_MAX_CHAR_NUM=20
## Show help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
