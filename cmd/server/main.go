package main

import (
	"github.com/Advanced-Memory-Analytics/process-exporter/api"
	"log"
)

func main() {
	err := api.StartServer()
	if err != nil {
		log.Fatalf("Server failed to start with error: %v.", err)
	}
	log.Println("Server stopped.")
}
