package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	jsonvalidator "github.com/go-playground/validator/v10"
	"github.com/nopilei/events/src/schema"
	"github.com/nopilei/events/src/transport/kafka"
	"github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

var validator *jsonvalidator.Validate = jsonvalidator.New()
var producer *kafka.KafkaProducer
var (
    httpRequests = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "path", "status"},
    )
)

func writeJSONError(w http.ResponseWriter, msg string, err error) {
	fmt.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ErrorResponse{Error: msg})
}

func sendEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	event := &schema.BaseEventData{}
	if err := json.NewDecoder(r.Body).Decode(event); err != nil{
		writeJSONError(w, "Invalid request body", err)
		return
	}

	eventFactory, ok := schema.EventsRegistry[event.EventType]
	if !ok {
		writeJSONError(w, "Invalid event_type.", errors.New("Unknown event type"))
		return
	}

	eventInstance := eventFactory()
	if err := json.Unmarshal(event.Data, eventInstance); err != nil{
		fmt.Printf("%v", err)
		writeJSONError(w, "Invalid event data.", err)
		return
	}

	if err := validator.Struct(eventInstance); err != nil {
		writeJSONError(w, "Invalid event data.", err)
		return
	}
	if producer == nil{
		writeJSONError(w, "Producer wasn't initialized", errors.New("Producer wasn't initialized"))
		return
	}
	if err := producer.Send(r.Context(), event, 5 * time.Second); err != nil{
		writeJSONError(w, "Message wasn't send.", err)
		return
	}
	
	httpRequests.WithLabelValues(r.Method, r.Pattern, "200").Inc()
	fmt.Fprintln(w, "Event received:", event)

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{status: "healthy"}`))
}

func main() {
	prometheus.MustRegister(httpRequests)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8000"
	}
	_producer, err := kafka.New()
	if err != nil{
		fmt.Print(err)
		return
	}
	producer = _producer

	mux := http.NewServeMux()
	mux.HandleFunc("POST /events", sendEvent)
	mux.HandleFunc("GET /health", healthCheck)
	mux.Handle("GET /metrics", promhttp.Handler())

	fmt.Println("starting server at", port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil{
		fmt.Printf("%v", err)
	}
}
