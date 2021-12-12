package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// This will only ever go up
var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "prom_go_processed_ops_total",
		Help: "The total number of processed events",
	})
)

// This can go up or down
var (
	opsInProgress = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "prom_go_progress_total",
		Help: "The total number of processed events",
	})
)

var histogram = promauto.NewHistogram(prometheus.HistogramOpts{
	Name:    "random_numbers",
	Help:    "A histogram of normally distributed random numbers.",
	Buckets: prometheus.LinearBuckets(-3, .1, 61),
})

func Random() {
	for {
		histogram.Observe(rand.NormFloat64())
	}
}

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			if rand.Int()%2 == 0 {
				opsInProgress.Inc()
			} else {
				opsInProgress.Dec()
			}
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	recordMetrics()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
