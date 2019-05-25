package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	configenv  EnvConfig
	configjson JSONConfig
)

// InitializeEnv will pull all the details from the env and put them in a struct
func InitializeEnv() {
	// Load .env
	godotenv.Load()

	// Define config structure
	configenv = EnvConfig{}

	if err := envconfig.Process("app", &configenv); err != nil {
		log.Println("Error trying to parse the env file, continuting anyway", err)
	}

	log.Println("CONFIG > ENV Initialized")
}

// InitializeJSON will pull all the details from the json and put them in a struct
func InitializeJSON() {
	// Read JSON from config file
	read, err := os.Open("./config.json")
	if err != nil {
		log.Fatalln("Error trying to parse the env file, stopping process", err)
		return
	}

	// Close the file
	defer read.Close()

	// Create new decoder
	decoder := json.NewDecoder(read)

	// Define config structure
	configjson = JSONConfig{}

	// Decode into the structure
	if err := decoder.Decode(&configjson); err != nil {
		log.Fatalln("Failed to decode json config, stopping process", err)
		panic(err)
	}

	log.Println("CONFIG > JSON Initialized")
}

// ENV will return the config
func ENV() EnvConfig {
	return configenv
}

// JSON will return the json config
func JSON() JSONConfig {
	return configjson
}

// FetchConfigForRepo to fetch the config for a specified repo
func FetchConfigForRepo(repo string) (*Repo, error) {
	// Define repo struct
	item := Repo{}

	// Loop over config elements
	for i := range configjson {
		// Check if repo is in config
		if configjson[i].Repo == strings.ToLower(repo) {
			// if it is set it
			item = configjson[i]
			break
		}
	}

	if item.Repo == "" {
		return &item, errors.New("not_found")
	}

	// Return item
	return &item, nil
}
