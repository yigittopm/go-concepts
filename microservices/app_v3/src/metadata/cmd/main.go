package main

import (
	"flag"
	"log"
)

const serviceName = "metadata"

func main() {
	var port int

	flag.IntVar(&port, "port", 8081, "APU handler port")
	flag.Parse()

	log.Printf("Starting the metadata service on port %d", port)

}
