package main

import (
	"flag"
	"log"

	"github.com/sebsprenger/chatterschool/client"
)

var (
	ip   = flag.String("ip", "localhost", "server ip")
	port = flag.String("port", "9001", "server port")
	name = flag.String("name", "nobody", "name used for chat")
)

func main() {
	flag.Parse()

	chatbot := NewChatbot(*name)

	client := client.NewChatClient()
	err := client.Connect(*ip, *port)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect()

	client.ReceiveChatMessagenOn(chatbot)
	chatbot.SendChatMessagesTo(&client)

}
