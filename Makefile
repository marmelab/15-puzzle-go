SIZE ?= 4

NAME := 15-puzzle-go
DOCKER := docker run -it --rm -v "${PWD}/src:/src" $(NAME)
DOCKER_WEBSERVER := docker run -it --rm -v "${PWD}/src:/src" -p 2000:2000 $(NAME)

help: ## Print all commands (default)
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build-docker: ## Build the docker
	docker build -t $(NAME) .

install: build-docker ## Build the docker and install the dependencies
	$(DOCKER) go get -u github.com/nsf/termbox-go
	$(DOCKER) go get -u github.com/gorilla/mux

run: ## Run the 15-puzzle game with the env variable SIZE as parameter
	$(DOCKER) go run main.go --size=$(SIZE)

run-server: ## Run the 15-puzzle webserver at port 2000
	$(DOCKER_WEBSERVER) go run main-server.go

test: ## Run all tests
	$(DOCKER) go test -v ./...

lint: ## Run the gofmt linter
	$(DOCKER) gofmt -w .
