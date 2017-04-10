package main

import (
	"fmt"
	"os"
	"strings"
	"bytes"

	"github.com/nlopes/slack"
)

func handleRTM(rtm *slack.RTM) {
	// Allow goroutines to manage the connection(s)
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			// Handle events
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter: ", ev.ConnectionCount)

			case *slack.MessageEvent:
				// Valid event
				fmt.Printf("Message: %v\n", ev)
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				if ev.User != info.User.ID &&
					strings.HasPrefix(ev.Text, prefix) {
					respond(rtm, ev, prefix)
				}

			case *slack.RTMError:
				fmt.Println("Error with rtm: ", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Println("Invalid credentials")
				os.Exit(1)

			default:
				// Do nothing
			}
		}
	}
}

func query(value string, param string) string {

	var buffer bytes.Buffer
	isComplex := true
	if param == "" {
		isComplex = false
	}

	switch value {
		case "list"
			if isComplex {
				buffer.WriteString("List info about: ")
				buffer.WriteString(param)
			} else {
				buffer.WriteString("List all repo's")
			}

		case "add":
			if isComplex {
				buffer.WriteString("Add: ")
				buffer.WriteString(param)
			} else {
				buffer.WriteString("ERROR: need repo uri")
			}

		case "remove":
			if isComplex {
				buffer.WriteString("Remove: ")
				buffer.WriteString(param)
			} else {
				buffer.WriteString("Error: need repo name")
			}

		case "activate":
			if isComplex {
				buffer.WriteString("Activate: ")
				buffer.WriteString(param)
			} else {
				buffer.WriteString("Error: need repo name")
			}

		case "deactivate":
			if isComplex {
				buffer.WriteString("Deactivate: ")
				buffer.WriteString(param)
			} else {
				buffer.WriteString("Error: need repo name")
			}
	}
	return buffer.String()
}

func respond(rtm *slack.RTM, msg *slack.MessageEvent, prefix string) {

	value := msg.Text
	value = strings.TrimPrefix(value, prefix)
	value = strings.TrimSpace(value)
	value = strings.ToLower(value)
	param := ""

	if strings.Contains(value, " ") {
		parts := strings.SplitN(value, " ", 2)
		value = parts[0]
		param = parts[1]
	}

	rtm.SendMessage(rtm.NewOutgoingMessage(query(value, param), msg.Channel))
}

func main() {
	// Add SLACK_TOKEN to computer's environment vars
	// ie: export SLACK_TOKEN="your_slack_bot_token"
	api := slack.New(os.Getenv("SLACK_TOKEN"))
	// api.SetDebug(true) // To get debug info uncomment this (set to false by default)

	// Slack's Real Time Message api
	handleRTM(api.NewRTM())
}
