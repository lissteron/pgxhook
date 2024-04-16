USER_ID := $(shell id -u):$(shell id -g)
DOCKER_COMPOSE_RUN ?= docker-compose
TEST_CMD ?= go clean -testcache && go test -race -v -tags=integration ./...

.PHONY: lint
lint: ## Run linter
	${DOCKER_COMPOSE_RUN} run --rm linter /bin/sh -c "golangci-lint run ./... -c .golangci.yml -v"

test: ## Run tests
	${DOCKER_COMPOSE_RUN} run --rm app /bin/sh -c "${TEST_CMD}"
	${DOCKER_COMPOSE_RUN} down

.PHONY: down
down: ## Down infra
	${DOCKER_COMPOSE_RUN} down --volumes

.PHONY: clean
clean: ## Remove all generated code
	rm -rf gen/

.PHONY: regen
regen: clean gen ## Regenerate all

.PHONY: gen
gen: ## Generate mocks
	mkdir -p ./gen/mocks
	${DOCKER_COMPOSE_RUN} run --rm mockery /bin/sh -c "mockery && chown -R $(USER_ID) gen/"
