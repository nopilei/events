package api

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
    HttpRequests = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "path", "status"},
    )
    HttpDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "http_duration_seconds",
            Help: "Duration of HTTP requests in seconds",
            Buckets: []float64{0.05, 0.1, 0.5, 1, 5, 10},
        },

        []string{"method", "path", "status"},
    )
)