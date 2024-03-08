APP_NAME=GlobalHitss Backend Challenge
VERSION?=0.0.1

GOOS=linux
GOARCH=amd64
CGO_ENABLED=0
GO111MODULE=auto
GOVERSION=1.21.1

USER_API_IMG_NAME=userapi
BIN_DIR=./bin

ARGS=$(filter-out $@,$(MAKECMDGOALS))

%:
	@:

.YELLOW := $(shell tput -Txterm setaf 3)
.RESET  := $(shell tput -Txterm sgr0)

.DEFAULT_GOAL := help

.PHONY: help

help:
	@echo "${APP_NAME} - v${VERSION}\n"
	@echo "Makefile targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo

# * GOLANG * #
## Clean built binaries directory
clean: 
	@echo "${.YELLOW}--- Go: clean ---${.RESET}"
	rm -rf $(BIN_DIR)

## Get dependencies
tidy: 
	@echo "${.YELLOW}--- Go: tidy ---${.RESET}"
	go mod tidy -compat=$(GOVERSION)

## Build binary
build: 
	@echo "${.YELLOW}--- Go: build ---${.RESET}"
	mkdir -p $(BIN_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) GO111MODULE=$(GO111MODULE) CGO_ENABLED=$(CGO_ENABLED) go build -o $(BIN_DIR) ./cmd/userapi

## Build local development binary
build-dev: 
	@echo "${.YELLOW}--- Go: build dev ---${.RESET}"
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR) ./cmd/userapi

## Run Go unit tests
tests: 
	@echo "${.YELLOW}--- Go: tests ---${.RESET}"
	go test -v -race ./...

## Run Go integration tests
tests-integration: 
	@echo "${.YELLOW}--- Go: tests ---${.RESET}"
	go test -tags e2e -v -race ./...

## Lint Go code
lint: 
	@echo "${.YELLOW}--- Go: lint ---${.RESET}"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run


# * DOCKER * #
## Build all docker images
docker-build: 
	@echo "${.YELLOW}--- Docker: build ---${.RESET}"
	docker build --build-arg SERVICE_NAME=$(USER_API_IMG_NAME) -t $(USER_API_IMG_NAME):latest .

## Build one docker image
docker-build-img: 
	@echo "${.YELLOW}--- Docker: build img ---${.RESET}"
	docker build --build-arg SERVICE_NAME=$(ARGS) -t $(ARGS):latest .

## Up docker-compose cluster
docker-up: docker-build 
	@echo "${.YELLOW}--- Docker: up ---${.RESET}"
	docker-compose up -d --build --remove-orphans 

## Down docker-compose cluster
docker-down: 
	@echo "${.YELLOW}--- Docker: down ---${.RESET}"
	docker-compose down --remove-orphans