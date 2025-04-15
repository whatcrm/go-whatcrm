package models

import "time"

type Message struct {
	ID               string     `json:"id,omitempty"`
	Content          string     `json:"content"`
	Type             string     `json:"type"`
	FileURL          string     `json:"file_url"`
	DialogID         string     `json:"dialog_id"`
	SenderID         string     `json:"sender_id"`
	CreatedAt        time.Time  `json:"timestamp"`
	IsEdited         bool       `json:"is_edited"`
	EditedAt         time.Time  `json:"edited_at"`
	Editable         bool       `json:"editable"`
	DeletedAt        time.Time  `json:"deleted_at"`
	DeletedStatus    string     `json:"deleted_status"`
	DeletedForUserID string     `json:"deleted_for"`
	Status           string     `json:"status"`
	Reactions        []Reaction `json:"reactions"`
	QuotedMessageID  string     `json:"quotedMsgId"`
	Timeout          int        `json:"timeout,omitempty"`
	Buttons          []Button   `json:"buttons,omitempty"`
}

type MessageInput struct {
	Body            string `json:"body"`
	ChatID          string `json:"chatId"`
	QuotedMessageID string `json:"quotedMsgId"`
	FileURL         string `json:"file_url"`
}

type MessageResponse struct {
	Data Message `json:"message"`
}
type Reaction struct {
	ChatID    string `json:"chat_id"`
	MessageID string `json:"message_id"`
	Reaction  string `json:"reaction"`
}

type Button struct {
	ID   int    `json:"id" bson:"id"`
	Body string `json:"body" bson:"body"`
}
