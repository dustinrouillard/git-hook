package discord

// Hook is the base discord hook json
type Hook struct {
	Embeds []Embed `json:"embeds,omitempty"`
}

// Embed defines the embed structure
type Embed struct {
	Title       string  `json:"title,omitempty"`
	URL         string  `json:"url,omitempty"`
	Description string  `json:"description,omitempty"`
	Color       int     `json:"color,omitempty"`
	Timestamp   string  `json:"timestamp,omitempty"`
	Author      Author  `json:"author,omitempty"`
	Fields      []Field `json:"fields,omitempty"`
	Footer      Footer  `json:"footer,omitempty"`
}

// Author will contain the embed author information
type Author struct {
	Name    string `json:"name,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

// Field defines an embed field structure
type Field struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

// Footer defines a footer for an embed
type Footer struct {
	Text string `json:"text,omitempty"`
}
