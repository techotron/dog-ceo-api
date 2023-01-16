REGISTRY=techotron/dog-ceo-api
APP=backend:$(VERSION)
IMAGE=golang:1.19 go
GOLANGCI_TAG=v1.50
ENVIRONMENT=dev

.PHONY: build
build: 
	docker build -t $(REGISTRY)/$(APP) .

.PHONY: run-db
run-db:
# Command to run db container for debugging of unit tests that rely on one
	docker-compose -f docker-compose.yaml -f docker-compose-debug.yaml up --build -d db

.PHONY: run-backend
run-backend:
	docker-compose -f docker-compose.yaml -f docker-compose-debug.yaml up --build -d backend db

.PHONY: up
up: run-backend

.PHONY: down
down:
	docker-compose -f docker-compose.yaml -f docker-compose-debug.yaml down -v

# https://docs.docker.com/compose/reference/up/
# Specify --exit-code-from test here to pass down the exit code to the calling process (docker-compose up). 
#   When this happens, the docker-compose up returns an exit code 1 which will attempt to stop all containers because it
#   "implies --abort-on-container-exit". Then the make down runs which results in a second SIGTERM getting sent, meaning
#   the resulting exit code is 2 - because: 
#     "If SIGINT or SIGTERM is sent again during this shutdown phase, the running containers are killed, and the exit code is 2.
# So if make test command exits with code 2, that means tests failed and make down command triggered an additional SIGINT signal 
#   which returned the code 2 instead of code 1 exit code
.PHONY: all-tests
all-tests:
	docker-compose up --exit-code-from test test db

.PHONY: test
test: all-tests down

.PHONY: lint
lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:$(GOLANGCI_TAG) golangci-lint run

.PHONY: publish
publish: docker-login	docker-push

.PHONY: coverage
coverage: test coverage-report

.PHONY: coverage-report
coverage-report: 
	go tool cover -html=c.out

.PHONY: coverage-ci
coverage-ci:
	docker run --rm -v $(shell pwd):/app -w /app $(IMAGE) tool cover -func=c.out -o report.out
