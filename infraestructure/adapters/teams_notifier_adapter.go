package adapters

import (
	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/jcastellanos/falcon/core/models"
	"log"
	"os"
)

type TeamsNotifierAdapter struct {

}

func NewTeamsNotifierAdapter() TeamsNotifierAdapter {
	return TeamsNotifierAdapter {}
}

func (a TeamsNotifierAdapter) Notify(monitor models.Monitor) (bool, error) {
	// Initialize a new Microsoft Teams client.
	mstClient := goteamsnotify.NewClient()

	// The title for message (first TextBlock element).
	msgTitle := "Alarma - Se ha presentado un error consultando el servicio"

	// Formatted message body.
	msgText := "Here are some examples of formatted stuff like " +
		"\n * this list itself  \n * **bold** \n * *italic* \n * ***bolditalic***"

	// Create message using provided formatted title and text.
	msg := goteamsnotify.NewMessageCard()
	msg.Text = msgText
	msg.Title = msgTitle

	// Send the message with default timeout/retry settings.
	if err := mstClient.Send(monitor.GuardChannel, msg); err != nil {
		log.Printf(
			"failed to send message: %v",
			err,
		)
		os.Exit(1)
	}
	return true, nil
}