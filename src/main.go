package main

import (
	"flag"
	"log"

	"github.com/markday1962/simple-http-server/pkg/server"
)

func main() {

	port := flag.String("port", "8080", "Port to listen to")
	flag.Parse()

	listeningPort := ":" + *port
	log.Println(listeningPort)

	httpServer := server.NewHTTPServer(listeningPort)

	if err := httpServer.Open(); err != nil {
		log.Fatal("could not open httpServer", err)
	}

}
