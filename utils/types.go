package utils

type Message struct {
	MessageId      int
	Text           string
	ReplyToMessage *Message
	ChatId         int64
}

type Request struct {
	Message *Message
}

type Response struct {
	ChatId int64
	Text   string
}

type Predicate interface {
}
