# 15-puzzle-go

A CLI tool to play the 15 puzzle game, in Go.

> See the [related article](https://marmelab.com/blog/2017/12/01/jeu-du-taquin-en-go.html) on the Marmelab blog

## Help

Print all available commands

```bash
make
```

## Build

### Install

Install the dependencies, compile the code and build the docker

```bash
make install
```

### Install prod

Install the dependencies, compile the code and build the docker in production mode

```bash
make install-prod
```

### Publish

Publish the docker in the docker hub [15-puzzle-api](https://hub.docker.com/r/luwangel/15-puzzle-api/).
_Note: be careful, you should be logged before!_

```bash
make publish
```

## Run the project

### Run the CLI game

Run the 15-puzzle game in dev or prod mode

```bash
make run
```

Note:

*   you can use the environment variable `ENV` to define the running env
*   you can use the environment variable `SIZE` to define the puzzle size

For example with a size equals to 5:

```bash
ENV=prod SIZE=5 make run
```

### Run the webserver

Run the 15-puzzle webserver at port 2000 in dev or prod mode

```bash
make run-server
```

Note:

*   you can use en environment variable `ENV` to define the running env
*   you can use en environment variable `PORT` to define the port (from 2000 to 3000)

```bash
ENV=prod PORT=2017 make run-server
```

## Contributing

### Test

Run all tests

```bash
make test
```

### Linter

Run the gofmt linter

```bash
make lint
```
