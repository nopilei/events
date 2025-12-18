package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/nopilei/events/src/schema"
	"github.com/nopilei/events/src/transport/kafka"
)


func SendEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	event := &schema.BaseEventData{}
	if err := json.NewDecoder(r.Body).Decode(event); err != nil{
		WriteJSONError(w, ErrWrongRequestBody, http.StatusBadRequest)
		return
	}

	eventFactory, ok := schema.EventsRegistry[event.EventType]
	if !ok {
		WriteJSONError(w, ErrWrongEventType, http.StatusBadRequest)
		return
	}

	eventInstance := eventFactory()
	if err := json.Unmarshal(event.Data, eventInstance); err != nil{
		WriteJSONError(w, ErrWrongEventData, http.StatusBadRequest)
		return
	}

	if err := Validate(eventInstance); err != nil {
		WriteJSONError(w, ErrWrongEventData, http.StatusBadRequest)
		return
	}
	
	if err := kafka.Send(r.Context(), event, 5 * time.Second); err != nil{
		WriteJSONError(w, err, http.StatusInternalServerError)
		return
	}
	
	HttpRequests.WithLabelValues(r.Method, r.Pattern, "200").Inc()
	fmt.Fprintln(w, "Event received:", event)

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{status: "healthy"}`))
}
