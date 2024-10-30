package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "80"

type ServerConfig struct{}

func main() {

	// var serverConfig Config
	serverConfig := ServerConfig{}

	log.Println("Start Main Function")

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: serverConfig.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
