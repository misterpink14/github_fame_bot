package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := slack.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
			case msg := <-rtm.IncomingEvents:
				fmt.Print("Event Received: "
				switch ev := msg.Data.(type) {
				case *slack.ConntectedEvent:
					fmt.Println("Conntection counter:", ev.ConntectionCount)
				case *slack.MessageEvent:
					fmt.Printf("Message: %v\n", ev)
					info := rtm.GetInfo()
					prefix := fmt.Sprintf("<@%s> ", info.User.ID)

					if ev.User != ev.User.ID && strings.HasPrefix(ev.Text, prefix)
						rtm.SendMessage(rtm.NewOutgoingMessage("What's up buddy!?!?", ev.Channel))
					}

				case *slack.RTMError:
					fmt.Printf("Error: %s\n", ev.Error())

				case *slack.InvalidAuthEvent:
					fmt.Printf("Invalid credentials")
					break Loop

				default:
					// Do nothing}}
}




