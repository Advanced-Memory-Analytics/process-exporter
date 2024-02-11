package api

import (
	"github.com/Advanced-Memory-Analytics/process-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/process"
	"log"
	"time"
)

var ProcessCollector = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "ProcessMetrics",
	Help: "Metrics for local processes",
}, []string{"metrics"})

func Collect() {
	processes, err := process.Processes()
	if err != nil {
		log.Fatalf("Failed to collect processes with error: %v", err)
	}
	for _, process := range processes {
		name, _ := process.Name()
		cpuUsage, _ := process.CPUPercent()
		ProcessCollector.WithLabelValues(name).Set(cpuUsage)
	}
}

func CollectWithInterval() {
	ticker := time.NewTicker(config.INTERVAL)
	go func() {
		for range ticker.C {
			Collect()
		}
	}()
}
