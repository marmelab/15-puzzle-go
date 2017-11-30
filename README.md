# 15-puzzle-go

A CLI tool to play the 15 puzzle game, in Go.

## Run the project

### Install

Install the dependencies, compile the code and build the docker

``` bash
make install
```

### Run the game

Run the 15-puzzle game

``` bash
make run
```

Note: you can use en environment variable `SIZE` to define the puzzle size

For example with a size equals to 5:

``` bash
SIZE=5 make run
```

### Run the webserver

Run the 15-puzzle webserver at port 2000

``` bash
make run-server
```

Note: you can use en environment variable `PORT` to define the port (from 2000 to 3000)

``` bash
PORT=2017 make run
```

## Contributing

### Help

Print all available commands

``` bash
make
```

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
