package api

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/process"
	"log"
)

type processCollector struct {
	cpuUsage *prometheus.Desc
	memUsage *prometheus.Desc
	idle     *prometheus.Desc
}

func NewProcessCollector() *processCollector {
	return &processCollector{
		cpuUsage: prometheus.NewDesc(
			"cpu_usage",
			"Percentage of CPU used by a process.",
			[]string{"process"},
			nil,
		),
		memUsage: prometheus.NewDesc(
			"memory_usage",
			"Percentage of RAM used by a process.",
			[]string{"process"},
			nil,
		),
		idle: prometheus.NewDesc(
			"is_idle",
			"Whether a process is idle or running.",
			[]string{"process"},
			nil,
		),
	}
}

func (collector *processCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.cpuUsage
	ch <- collector.memUsage
	ch <- collector.idle
}

func (collector *processCollector) Collect(ch chan<- prometheus.Metric) {
	processes, err := process.Processes()
	if err != nil {
		log.Fatalf("Failed to collect processes with error: %v", err)
	}
	for _, process := range processes {
		idle := 0.0

		pid := process.Pid

		name, err := process.Name()
		if err != nil {
			continue
		}

		cpuUsage, err := process.CPUPercent()
		if err != nil {
			continue
		}

		memUsage, err := process.MemoryPercent()
		if err != nil {
			continue
		}

		running, err := process.IsRunning()
		if err != nil {
			continue
		}
		if running {
			idle = 1.0
		}

		ch <- prometheus.MustNewConstMetric(collector.cpuUsage, prometheus.GaugeValue, cpuUsage, fmt.Sprintf("PID:%d_%s_%s", pid, name, "cpu"))
		ch <- prometheus.MustNewConstMetric(collector.memUsage, prometheus.GaugeValue, float64(memUsage), fmt.Sprintf("PID:%d_%s_%s", pid, name, "mem"))
		ch <- prometheus.MustNewConstMetric(collector.idle, prometheus.GaugeValue, idle, fmt.Sprintf("PID:%d_%s_%s", pid, name, "running"))
	}

}
