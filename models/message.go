package models

import "time"

type Message struct {
	ID               string     `json:"id,omitempty" bson:"_id,omitempty"`
	Content          string     `json:"content" bson:"content"`
	Type             string     `json:"type" bson:"type"`
	FileURL          string     `json:"file_url" bson:"file_url"`
	DialogID         string     `json:"dialog_id" bson:"dialog_id"`
	SenderID         string     `json:"sender_id" bson:"sender_id"`
	CreatedAt        time.Time  `json:"timestamp" bson:"created_at"`
	IsEdited         bool       `json:"is_edited" bson:"is_edited"`
	EditedAt         time.Time  `json:"edited_at" bson:"edited_at"`
	Editable         bool       `json:"editable" bson:"editable"`
	DeletedAt        time.Time  `json:"deleted_at" bson:"deleted_at"`
	DeletedStatus    string     `json:"deleted_status" bson:"deleted_status"`
	DeletedForUserID string     `json:"deleted_for" bson:"deleted_for"`
	Status           string     `json:"status" bson:"status"`
	Reactions        []Reaction `json:"reactions" bson:"reactions"`
	QuotedMessageID  string     `json:"quotedMsgId" bson:"quotedMsgId"`
	Timeout          int        `json:"timeout,omitempty" bson:"timeout,omitempty"`
	Buttons          []Button   `json:"buttons,omitempty" bson:"buttons,omitempty"`
}

type Reaction struct {
	MessageID string `json:"message_id" bson:"message_id"`
	UserID    string `json:"user_id" bson:"user_id"`
	Symbol    string `json:"symbol" bson:"symbol"`
}

type Button struct {
	ID   int    `json:"id" bson:"id"`
	Body string `json:"body" bson:"body"`
}
