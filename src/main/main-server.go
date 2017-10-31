package main

import (
	"fmt"
	"webserver"
)

const DEFAULT_WEBSERVER_PORT int = 2000

func main() {
	fmt.Printf("Starting the server at port %d\n", DEFAULT_WEBSERVER_PORT)
	webserver.Server(DEFAULT_WEBSERVER_PORT)
}
