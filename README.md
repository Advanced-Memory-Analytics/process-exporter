# process-exporter
Experimental metric exporter for processes running locally.

## Running
Server opens on port 9098 by default. Options might be added later to configure this
the via command line.  Visit localhost:9098/metrics to see metrics.  You can build an executable or run
the server using ```go run cmd/server/main.go``` (assuming you are in the root of the project).

## Configure with Prometheus
A simple configuration may look something like this:

```
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: "node"
    static_configs:
    - targets: ["localhost:9098"]

```
## Configure with Grafana
Start your prometheus server with the config file set to target your process-exporter server.
From here you can add a data source that gathers metrics from your prometheus server.

## Acknowledgements
This project relies heavily on an open source library to get process information.
That open source project can be found here: [gopsutil](https://github.com/shirou/gopsutil?tab=readme-ov-file)

This project also uses Prometheus' Golang client.
That can be found here: [prometheus/client_golang](https://github.com/prometheus/client_golang)