# Build variables
LATEST_COMMIT := $$(git rev-parse HEAD)

.PHONY: help build docker up down logs kup kdown klogs test

help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
%:
	@:

build: ## Build the app
	@go clean
	@CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build \
	-a -installsuffix nocgo \
	-ldflags "-X main.buildTag=`date -u +%Y%m%d.%H%M%S`-$(LATEST_COMMIT)" \
	-o ./sta .

docker: ## Build docker image
	@docker build -t sta:latest .
	@go clean

up: ## Run docker container
	@docker run -d --name=sta-test -e APP_PORT=8080 -p 8080:8080 sta:latest

down: ## Down docker container and remove it with the image
	@docker stop sta-test && docker rm sta-test && docker rmi sta:latest

logs: ## Show logs of the sta-test docker container
	@docker logs sta-test

test: ## Run all tests
	@go test ./interactor
	@go test ./interactor/h
	@go test ./interactor/k

