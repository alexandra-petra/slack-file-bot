package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	// setting up the environment

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5015548833060-5015940101972-rctctFJ35mUap2uR7vhYUS00") // xoxb-5015548833060-5015940101972-rctctFJ35mUap2uR7vhYUS00
	os.Setenv("CHANNEL_ID", "D04VCTM79MM")                                                    // member ID: U050FTN2ZUL  C050RLJE933
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")} // my channel
	fileArr := []string{"tenses.pdf"}               // file slice

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{ // func from slack
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil { // there is en error
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}
