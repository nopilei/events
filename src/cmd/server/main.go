package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nopilei/events/src/transport/kafka"
	"github.com/nopilei/events/src/api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	prometheus.MustRegister(api.HttpRequests)

	err := kafka.InitProducer()
	if err != nil {
		fmt.Print(err)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /events", api.SendEvent)
	mux.HandleFunc("GET /health", api.HealthCheck)
	mux.Handle("GET /metrics", promhttp.Handler())

	port := os.Getenv("API_PORT")
	fmt.Println("starting server at", port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
