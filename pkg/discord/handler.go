package discord

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// SendHook will send the supplied payload to the specified discord webhook
func SendHook(payload Hook, url string) {
	// Martshaling incoming payload
	body, jsonErr := json.Marshal(payload)
	if jsonErr != nil {
		log.Println("Error in sending Discord hook", jsonErr)
	}

	// Create http client
	client := http.Client{}

	// Set request information for discord webhook
	request, discordErr := http.NewRequest("POST", url, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	if discordErr != nil {
		log.Println("Error in sending Discord hook", discordErr)
		return
	}

	// Do client request
	res, doErr := client.Do(request)
	if doErr != nil {
		log.Println("Error in sending Discord hook", doErr)
		return
	}
	defer res.Body.Close()
}
