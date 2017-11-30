package main

import (
	"flag"
	"fmt"
	"webserver"
)

const DEFAULT_WEBSERVER_PORT int = 2000

func getPort() int {
	var port int
	flag.IntVar(&port, "port", DEFAULT_WEBSERVER_PORT, "Port on which application should listen")
	flag.Parse()
	return port
}

func main() {
	port := getPort()
	fmt.Printf("Starting the server at port %d\n", port)
	webserver.Server(port)
}
