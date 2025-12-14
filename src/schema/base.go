package schema

import "encoding/json"

type BaseEventData struct {
	EventType string `json:"event_type"`
	Data      json.RawMessage `json:"data"`
}
