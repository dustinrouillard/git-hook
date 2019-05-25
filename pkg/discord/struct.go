package discord

type Hook struct {
	Embeds []Embed `json:"embeds,omitempty"`
}

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

type Author struct {
	Name    string `json:"name,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type Field struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

type Footer struct {
	Text string `json:"text,omitempty"`
}
