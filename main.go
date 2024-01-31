package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "SLACK_BOT_TOKEN")
	os.Setenv("CHANNEL_ID", "C06FWKUUP2B")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	chnarr := []string{os.Getenv("CHANNEL_ID")}
	filearr := []string{"README.md"}
	for i := 0; i < len(filearr); i++ {
		params := slack.FileUploadParameters{
			Channels: chnarr,
			File:     filearr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("error:%v", err)
		}
		fmt.Printf("Name :%v URl: %v", file.Name, file.URL)
	}

}
