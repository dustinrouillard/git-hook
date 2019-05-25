package github

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dustin/go-humanize"

	"dev.tetrabyte/git-hook/internal/pkg/config"
	"dev.tetrabyte/git-hook/internal/pkg/utilities"
	"dev.tetrabyte/git-hook/pkg/discord"
)

// PingHandler will handle the ping event from github
func PingHandler(w http.ResponseWriter, r *http.Request) {
	// Define structure
	Data := PingEvent{}

	// Decode Body
	raw, decodeErr := ioutil.ReadAll(r.Body)
	if decodeErr != nil {
		log.Println("HTTP > Failed to decode body in github event", decodeErr)

		// Send failure
		utilities.DefaultResponse("internal_error", "Internal Server Error", 500, w)
		return
	}

	// Unmarshal body
	if jsonErr := json.Unmarshal([]byte(raw), &Data); jsonErr != nil {
		log.Println("HTTP > Failed to unmarshal body in github event", jsonErr)

		// Send failure
		utilities.DefaultResponse("internal_error", "Internal Server Error", 500, w)
		return
	}

	// Get Config for Repo
	Config, configErr := config.FetchConfigForRepo(Data.Repository.FullName)
	if configErr != nil {
		switch configErr.Error() {
		case "not_found":
			utilities.DefaultResponse("respository_not_supported", "Repository not supported", 400, w)
			break
		default:
			log.Println("Error getting config for repo", configErr)
			utilities.DefaultResponse("respository_not_supported", "Repository not supported", 400, w)
			break
		}

		return
	}

	// Verify Secret and Body Details
	verifyErr := Verify(Config.Secret, r.Header.Get("x-hub-signature"), raw)
	if verifyErr != nil {
		switch verifyErr.Error() {
		case "no_signature":
			utilities.DefaultResponse("missing_signature", "No signature provided", 400, w)
			break
		case "invalid_signature":
			utilities.DefaultResponse("missing_signature", "Invalid signature", 400, w)
			break
		default:
			log.Println("GIT > Failed to verify secret for push from "+Data.Repository.FullName, verifyErr)
			utilities.DefaultResponse("internal_error", "Internal Server Error", 500, w)
			break
		}

		return
	}

	// Log about it
	log.Println("GIT > Webhoook for Repository " + Data.Repository.FullName + " initialized")

	// Send Response
	utilities.DefaultResponse("success", "Pong", 200, w)
}

// PushHandler will handle the push event from github
func PushHandler(w http.ResponseWriter, r *http.Request) {
	// Define structure
	Data := PushEvent{}

	// Decode Body
	raw, decodeErr := ioutil.ReadAll(r.Body)
	if decodeErr != nil {
		log.Println("HTTP > Failed to decode body in github event", decodeErr)

		// Send failure
		utilities.DefaultResponse("internal_error", "Internal Server Error", 500, w)
		return
	}

	// Unmarshal body
	if jsonErr := json.Unmarshal([]byte(raw), &Data); jsonErr != nil {
		log.Println("HTTP > Failed to unmarshal body in github event", jsonErr)

		// Send failure
		utilities.DefaultResponse("internal_error", "Internal Server Error", 500, w)
		return
	}

	// Get Config for Repo
	Config, configErr := config.FetchConfigForRepo(Data.Repository.FullName)
	if configErr != nil {
		switch configErr.Error() {
		case "not_found":
			utilities.DefaultResponse("respository_not_supported", "Repository not supported", 400, w)
			break
		default:
			log.Println("Error getting config for repo", configErr)
			utilities.DefaultResponse("respository_not_supported", "Repository not supported", 400, w)
			break
		}

		return
	}

	// Verify Secret and Body Details
	verifyErr := Verify(Config.Secret, r.Header.Get("x-hub-signature"), raw)
	if verifyErr != nil {
		switch verifyErr.Error() {
		case "no_signature":
			utilities.DefaultResponse("missing_signature", "No signature provided", 400, w)
			break
		case "invalid_signature":
			utilities.DefaultResponse("missing_signature", "Invalid signature", 400, w)
			break
		default:
			log.Println("GIT > Failed to verify secret for push from "+Data.Repository.FullName, verifyErr)
			utilities.DefaultResponse("internal_error", "Internal Server Error", 500, w)
			break
		}

		return
	}

	// Log about it
	log.Println("GIT > New push hook from " + Data.Repository.FullName)

	// Ignore if commits length is 0
	if len(Data.Commits) < 1 {
		log.Println("GIT > Push has no commits not sending " + Data.Repository.FullName)
		return
	}

	// Commits text formatter
	CommitText := "commit"
	if len(Data.Commits) > 1 {
		CommitText = "commits"
	}

	// Commit length checker
	CommitsLength := len(Data.Commits)
	if len(Data.Commits) > 10 {
		CommitsLength = 10
	}

	// Description(CommitsDetails) text formatter
	var DescriptionBuffer bytes.Buffer
	for i := 0; i <= CommitsLength-1; i++ {
		log.Println(i)
		// Message Length Formatter
		CommitMessage := Data.Commits[i].Message
		if len(Data.Commits[i].Message) > 48 {
			CommitMessage = Data.Commits[i].Message[:45] + "..."
		}

		DescriptionBuffer.WriteString("[`" + Data.Commits[i].ID[:8] + "`](" + Data.Commits[i].URL + ") | " + CommitMessage + " - " + Data.Commits[i].Author.Name + "")
	}

	Description := DescriptionBuffer.String()

	// Append Special Text at the end if commits were removed due to length
	if len(Data.Commits) > 10 {
		Description = Description + "\nand " + humanize.Comma(int64(len(Data.Commits)-CommitsLength))
	}

	// Build payload
	Payload := discord.Embed{
		Title:       "[" + Data.Repository.Name + ":" + Data.Ref[11:] + "] | " + humanize.Comma(int64(len(Data.Commits))) + " new " + CommitText,
		URL:         Data.Repository.HTMLURL,
		Description: Description,
		Timestamp:   time.Now().Format(time.RFC3339),
		Footer: discord.Footer{
			Text: "Tetrabyte Git Hook",
		},
	}

	// Send Webhook Messages
	for i := range Config.Discord {
		discord.SendHook(discord.Hook{Embeds: []discord.Embed{Payload}}, Config.Discord[i])
	}

	// Send Response
	utilities.DefaultResponse("success", "Successfully received push hook", 200, w)
}
