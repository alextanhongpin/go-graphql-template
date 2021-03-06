ENV ?= development
-include .env.$(ENV)
-include .env
export

start: gen
	@go run cmd/server/{main,wire_gen}.go

# Include it at the bottom so that the commands sequence is respected.
include Makefile.db Makefile.dk

.PHONY: worker
worker:
	workwebui -redis="redis://${REDIS_HOST}:${REDIS_PORT}" -ns="question_worker" -listen=":5040"

test: # Runs unit testing without any infrastructure dependencies such as database.
	@#Do this to pipe to both file and stdout.
	@#go test -coverprofile cover.out -v ./... | tee unit.out
	go test -coverprofile cover.out -v ./...
	go tool cover -html=cover.out

# Alias for unit-testing.
ut: test

integration-test: # Run integration tests with the +build integration flag.
	go test -v -failfast -coverprofile cover.out -tags=integration ./...
	go tool cover -html=cover.out

# Alias for integration-testing.
it: integration-test


service-%:
	gen generate -t domain $*

gen: ## Generate packr migration bindata.
	@go generate ./...
