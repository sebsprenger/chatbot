package main

import (
	"fmt"

	"github.com/sebsprenger/chatbot/plugin"
	"github.com/sebsprenger/chatterschool/client"
	"github.com/sebsprenger/chatterschool/shared"
)

type Chatbot struct {
	messages chan shared.Message
	sender   string
	plugin   plugin.Bot
}

func NewChatbot(name string) Chatbot {
	return Chatbot{
		messages: make(chan shared.Message, 100),
		sender:   name,
		plugin:   plugin.Bot{},
	}
}

func (bot Chatbot) FormatMessage(msg shared.Message) {
	bot.messages <- msg
}

func (bot Chatbot) SendChatMessagesTo(client *client.ChatClient) {
	for inputMessage := range bot.messages {
		response := bot.plugin.Respond(inputMessage)

		if response == plugin.NOTHING {
			continue
		}
		outputMessage := bot.buildMessage(response)

		err := client.Send(outputMessage)
		if err != nil {
			fmt.Printf("Error while sending message: %s\n", err)
		}
	}
}

func (bot Chatbot) buildMessage(input string) shared.Message {
	msg := shared.Message{
		Text:   input,
		Sender: bot.sender,
	}
	return msg
}
