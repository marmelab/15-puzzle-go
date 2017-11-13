# 15-puzzle-go

A CLI tool to play the 15 puzzle game, in Go.

## Help

Print all available commands

``` bash
make
```

## Build

### Install

Install the dependencies, compile the code and build the docker

``` bash
make install
```

### Install prod

Install the dependencies, compile the code and build the docker in production mode

``` bash
make install-prod
```

### Publish

Publish the docker in the docker hub [15-puzzle-api](https://hub.docker.com/r/luwangel/15-puzzle-api/).
_Note: be careful, you should be logged before!_

``` bash
make publish
```

## Run the project

### Run the game

Run the 15-puzzle game in dev or prod mode

``` bash
make run
make run-prod
```

_Note: you can use en environment variable `SIZE` to define the puzzle size_

For example with a size equals to 5:

``` bash
SIZE=5 make run
SIZE=5 make run-prod
```

### Run the webserver

Run the 15-puzzle webserver at port 2000 in dev or prod mode

``` bash
make run-server
make run-server-prod
```

Note: you can use en environment variable `PORT` to define the port (from 2000 to 3000)

``` bash
PORT=2017 make run
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
