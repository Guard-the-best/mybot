package main

import (
	"github.com/Guard-the-best/mybot/utils"
	"github.com/mymmrac/telego"
)

func BuildMessage(inMessage *telego.Message) *utils.Message {
	var outMessage utils.Message

	if inMessage != nil {
		outMessage.MessageId = inMessage.MessageID
		outMessage.Text = inMessage.Text
		outMessage.ChatId = inMessage.Chat.ID
		outMessage.ReplyToMessage = BuildMessage(inMessage.ReplyToMessage)
	}

	return &outMessage
}

func BuildRequest(update telego.Update) utils.Request {
	var request utils.Request
	request.Message = BuildMessage(update.Message)
	return request
}
