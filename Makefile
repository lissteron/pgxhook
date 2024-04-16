USER_ID := $(shell id -u):$(shell id -g)
DOCKER_COMPOSE_RUN ?= docker-compose

TEST_CMD ?= go test -race -v -tags=integration ./...
WITH_COVER ?= -coverprofile=.cover/cover.out && go tool cover -html=.cover/cover.out -o .cover/cover.html

default: lint test

.cover/:
	mkdir $@

.PHONY: lint
lint: ## Run linter
	${DOCKER_COMPOSE_RUN} run --rm linter /bin/sh -c "golangci-lint run ./... -c .golangci.yml -v"

test: .cover/ ## Run tests
	${DOCKER_COMPOSE_RUN} run --rm app /bin/sh -c "${TEST_CMD} ${WITH_COVER}"
	${DOCKER_COMPOSE_RUN} down

.PHONY: down
down: ## Down infra
	${DOCKER_COMPOSE_RUN} down --volumes

.PHONY: clean
clean: ## Remove all generated code
	rm -rf gen/
	rm -rf .cover/

.PHONY: regen
regen: clean gen ## Regenerate all

.PHONY: gen
gen: ## Generate mocks
	mkdir -p ./gen/mocks
	${DOCKER_COMPOSE_RUN} run --rm mockery /bin/sh -c "mockery && chown -R $(USER_ID) gen/"
