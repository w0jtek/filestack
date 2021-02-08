package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// AcceptResponse defines the structure of the accept endpoint response
type AcceptResponse struct {
	HTTPCode int
	Message  string `json:"message"`
}

// NewAcceptResponse constructor
func NewAcceptResponse(httpCode int, message string) AcceptResponse {
	return AcceptResponse{
		HTTPCode: httpCode,
		Message:  message,
	}
}

// Render renders response
func (ar AcceptResponse) Render(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(ar.HTTPCode)
	bytes, _ := json.Marshal(ar)
	fmt.Fprintf(w, string(bytes))
}
