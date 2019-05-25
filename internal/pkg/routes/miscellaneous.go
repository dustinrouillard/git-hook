package routes

import "net/http"

// PingRoute used as a handler func will send back a json object to the client
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ "code": "ping", "message": "Ping!" }`))
}

// NotFound used as a handler func will send back a json object to the client
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{ "code": "endpoint_not_found", "message": "Endpoint not found" }`))
}
