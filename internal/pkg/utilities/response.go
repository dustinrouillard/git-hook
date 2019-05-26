package utilities

import (
	"encoding/json"
	"net/http"
)

// DefaultResponsePayload is the default response json structure
type DefaultResponsePayload struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// DefaultResponse will build a response based on the given params
func DefaultResponse(Code string, Message string, Status int, Writer http.ResponseWriter) {
	// Create response struct
	response := DefaultResponsePayload{
		Code:    Code,
		Message: Message,
	}

	// Marshal the struct
	jsonBody, marshalErr := json.Marshal(response)
	if marshalErr != nil {
		// If marshal fails say internal server error
		Writer.Header().Set("Content-Type", "application/json")
		Writer.WriteHeader(http.StatusInternalServerError)
		Writer.Write([]byte(`{ "code": "internal_error", "message": "Internal server error" }`))
		return
	}

	// Add json header
	Writer.Header().Set("Content-Type", "application/json")

	// Add status code
	Writer.WriteHeader(Status)

	// Write body
	Writer.Write([]byte(jsonBody))
}
