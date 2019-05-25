package routes

import (
	"log"
	"net/http"

	"dev.tetrabyte/git-hook/internal/pkg/github"
)

// Github used as a handler func will send back a json object to the client
func Github(w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get("x-github-event") {
	case "ping":
		github.PingHandler(w, r)
		break
	case "push":
		github.PushHandler(w, r)
		break
	default:
		log.Println("GIT > Got unsupported event " + r.Header.Get("x-github-event"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"code": "unsupported_event", "message": "Invalid Event Type"}`))
		break
	}
}
