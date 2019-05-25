package rest

import (
	"log"
	"net/http"

	"dev.tetrabyte/git-hook/internal/pkg/config"
	"dev.tetrabyte/git-hook/internal/pkg/git"
	"dev.tetrabyte/git-hook/internal/pkg/routes"

	"github.com/go-chi/chi"
)

// InitializeHttp will initialize the http server and setup the routes
func Initialize() {
	// Create New Router
	Router := chi.NewRouter()

	// Initialize Routes
	Router.Route("/github", git.Github)

	// Catch all 404 Route
	Router.NotFound(routes.NotFound)

	// Log about initialized
	log.Println("HTTP > Initialized on " + config.ENV().Port)

	// Start http server using go func to allow the function to continue
	log.Fatalln(http.ListenAndServe(":"+config.ENV().Port, Router))
}
