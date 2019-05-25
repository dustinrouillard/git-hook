package main

import (
	"dev.tetrabyte/git-hook/internal/app/rest"
	"dev.tetrabyte/git-hook/internal/pkg/config"
)

// Main Function
func main() {
	// Initialize Config Env
	config.InitializeEnv()

	// Initialize Config Json
	config.InitializeJSON()

	// Initialize HTTP
	rest.Initialize()
}
