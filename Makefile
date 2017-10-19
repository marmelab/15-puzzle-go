DOCKER := docker run -it --rm -v "${PWD}/src:/src" 15-puzzle-go

help: ## Print all commands (default)
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install: ## Install the dependencies, compile the code and build the docker
	docker build -t 15-puzzle-go .

run: ## Run the 15-puzzle game
	$(DOCKER) go run main.go

test: ## Run all tests
	$(DOCKER) go test -v ./...

lint: ## Run the gofmt linter
	$(DOCKER) gofmt -w .
