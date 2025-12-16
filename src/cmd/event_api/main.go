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
)

type ErrorResponse struct {
	Error string `json:"error"`
}

var validator *jsonvalidator.Validate = jsonvalidator.New()
var producer *kafka.KafkaProducer

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
	if err := producer.Send(r.Context(), eventInstance, time.Second); err != nil{
		writeJSONError(w, "Message wasn't send.", err)
		return
	}
	
	fmt.Fprintln(w, "Event received:", eventInstance)

}

func main() {
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

	fmt.Println("starting server at", port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil{
		fmt.Printf("%v", err)
	}
}
