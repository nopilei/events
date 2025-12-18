package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	jsonvalidator "github.com/go-playground/validator/v10"
	"github.com/nopilei/events/src/schema"
)

var (
	ErrWrongRequestBody = errors.New("wrong request body")
	ErrWrongEventType = errors.New("wrong event_type")
	ErrWrongEventData = errors.New("invalid event data")

	validator *jsonvalidator.Validate = jsonvalidator.New()
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteJSONError(w http.ResponseWriter, err error, statusCode int) {
	fmt.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
}

func Validate(event schema.Event) error {
	return validator.Struct(event)
}