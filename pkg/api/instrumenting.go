package api

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// RecordMetrics ...
func RecordMetrics() {
	go func() {
		for {
			OpsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

// OpsProcessed ...
var (
	OpsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processes_ops_total",
		Help: "The total number of processed events",
	})
)
