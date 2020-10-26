-include .env
export

start: gen
	@go run main.go

pkg:
	@pkger -include /schema

# Include it at the bottom so that the commands sequence is respected.
include pkg/database/Makefile

.PHONY: worker
worker:
	workwebui -redis="redis://${REDIS_HOST}:${REDIS_PORT}" -ns="question_worker" -listen=":5040"

test:
	@#Do this to pipe to both file and stdout.
	@#go test -coverprofile cover.out -v ./... | tee unit.out
	go test -coverprofile cover.out -v ./...
	go tool cover -html=cover.out

# Alias for unit-testing.
ut: test

integration-test:
	go test -v -failfast -coverprofile cover.out -tags=integration ./...
	go tool cover -html=cover.out

# Alias for integration-testing.
it: integration-test


up:
	@docker-compose up -d

down:
	@docker-compose down
