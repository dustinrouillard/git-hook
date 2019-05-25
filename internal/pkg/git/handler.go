package git

import (
	"dev.tetrabyte/git-hook/internal/pkg/routes"
	"github.com/go-chi/chi"
)

// Github handles the routes for /github
func Github(Router chi.Router) {
	// Root
	Router.Route("/", func(Root chi.Router) {
		// Ping on Get
		Root.Get("/", routes.Ping)

		// Handle Event
		Root.Post("/", routes.Github)
	})
}
