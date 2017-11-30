package main

import (
	"fmt"
	"flag"
	"webserver"
)

const DEFAULT_WEBSERVER_PORT int = 2000

func getPort() int {
	var port int
	flag.IntVar(&port, "port", DEFAULT_WEBSERVER_PORT, "an int")
	flag.Parse()
	if port > 2000 && port < 3000 {
		return port
	}
	return DEFAULT_WEBSERVER_PORT
}

func main() {
	port := getPort()
	fmt.Printf("Starting the server at port %d\n", port)
	webserver.Server(port)
}
