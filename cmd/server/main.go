package main

import (
	"github.com/Advanced-Memory-Analytics/process-exporter/api"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"os"
)

func init() {
	prometheus.MustRegister(api.NewProcessCollector())
}

func main() {
	log.SetOutput(os.Stdout)
	err := api.StartServer()
	if err != nil {
		log.Fatalf("Server failed to start with error: %v.", err)
	}
}
