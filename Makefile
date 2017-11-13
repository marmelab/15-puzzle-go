SIZE ?= 4
PORT ?= 2000

CONTAINER_NAME := 15-puzzle-go
DOCKER := docker run -it --rm -v "${PWD}/src:/src" $(CONTAINER_NAME)
DOCKER_WEBSERVER := docker run -it --rm -v "${PWD}/src:/src" -p $(PORT):$(PORT) $(CONTAINER_NAME)

DOCKERFILE_NAME_PROD := DockerfileProd.docker
CONTAINER_NAME_PROD := 15-puzzle-api
REPOSITORY_NAME_PROD := luwangel
DOCKER_PROD := docker run -it --rm $(CONTAINER_NAME_PROD)
DOCKER_WEBSERVER_PROD := docker run -it --rm -p $(PORT):$(PORT) $(CONTAINER_NAME_PROD)

help: ## Print all commands (default)
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

####### BUILD #######

build-docker: ## Build the dev docker
	docker build -t $(CONTAINER_NAME) .

install: build-docker ## Build the dev docker (alias for `build-docker`)

build-docker-prod: ## Build the prod docker
	docker build -f $(DOCKERFILE_NAME_PROD) -t $(REPOSITORY_NAME_PROD)/$(CONTAINER_NAME_PROD) .

install-prod: build-docker-prod ## Build the prod docker (alias for `build-docker-prod`)

publish: build-docker-prod ## Publish the docker in the dockerhub. Be careful, you should be logged before!
	docker push $(REPOSITORY_NAME_PROD)/$(CONTAINER_NAME_PROD)

####### RUN #######

run: ## Run the 15-puzzle game with the env variable SIZE as parameter
	$(DOCKER) go run main/main.go --size=$(SIZE)

run-server: ## Run the 15-puzzle webserver at port (default: 2000)
	$(DOCKER_WEBSERVER) go run main/main-server.go  --port=$(PORT)

####### DEV #######

run-prod: ## Run the 15-puzzle game with the env variable SIZE as parameter
	$(DOCKER) go run main/main.go --size=$(SIZE)

run-server-prod: ## Run the 15-puzzle webserver at port 2000
	$(DOCKER_WEBSERVER) go run main/main-server.go

####### DEV #######

test: ## Run all tests
	$(DOCKER) go test -v ./...

lint: ## Run the gofmt linter
	$(DOCKER) gofmt -w .
