package api

import "github.com/prometheus/client_golang/prometheus"

var ProcessCollector = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "ProcessMetrics",
	Help: "Metrics for local processes",
}, []string{"metrics"})
