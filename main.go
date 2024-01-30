package main

import (
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6555216986818-6548730551510-nHqT3m7OSqyWDYE2VlCMmgPP")
	os.Setenv("CHANNEL_ID", "C06FWKUUP2B")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	chnarr := []string{os.Getenv("CHANNEL_ID")}
	filearr := []string{""}

}
