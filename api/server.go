package api

import (
	"fmt"
	"github.com/Advanced-Memory-Analytics/process-exporter/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func StartServer() error {
	port := config.WEB_PORT
	http.Handle("/metrics", promhttp.Handler())

	log.Printf("Server: process-exporter started at %d.", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
