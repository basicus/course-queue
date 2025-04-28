package models

import "time"

type Message struct {
	Timestamp   time.Time `json:"timestamp"`
	MessageId   string    `json:"message_id"`
	MessageText string    `json:"message_text"`
}
type MessageReply struct {
	MessageId     string `json:"message_id"`
	MessageStatus int    `json:"message_status"`
	MessageReply  string `json:"message_reply"`
}
