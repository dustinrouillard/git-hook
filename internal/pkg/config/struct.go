package config

// EnvConfig Structure
type EnvConfig struct {
	Port string `envconfig:"PORT" default:"9090"`
}

// JSONConfig Structure
type JSONConfig []Repo

// Repo for json config
type Repo struct {
	Secret  string          `json:"secret"`
	Repo    string          `json:"repo"`
	Events  map[string]bool `json:"events"`
	Discord []string        `json:"discord"`
}
