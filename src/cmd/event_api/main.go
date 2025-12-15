package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	jsonvalidator "github.com/go-playground/validator/v10"
	"github.com/nopilei/events/src/schema"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

var validator *jsonvalidator.Validate = jsonvalidator.New()

func writeJSONError(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ErrorResponse{Error: msg})
}

func sendEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	event := &schema.BaseEventData{}
	err := json.NewDecoder(r.Body).Decode(event)
	if err != nil {
		writeJSONError(w, "Invalid request body")
		return
	}

	eventFactory, ok := schema.EventsRegistry[event.EventType]
	if !ok {
		writeJSONError(w, "Invalid event_type.")
		return
	}

	eventInstance := eventFactory()
	err = json.Unmarshal(event.Data, eventInstance)
	if err != nil {
		fmt.Printf("%v", err)
		writeJSONError(w, "Invalid event data.")
		return
	}

	err = validator.Struct(eventInstance)
	if err != nil {
		writeJSONError(w, "Invalid event data.")
		return
	}
	fmt.Fprintln(w, "Event received:", eventInstance)

}

func main() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8000"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("POST /events", sendEvent)

	fmt.Println("starting server at", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil{
		fmt.Printf("%v", err)
	}
}
