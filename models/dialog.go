package models

import "time"

type Dialog struct {
	ID             string    `json:"id,omitempty" bson:"_id,omitempty"`
	ChatID         string    `json:"chat_id" bson:"chat_id"`
	Guest          User      `json:"guest,omitempty" bson:"guest,omitempty"`
	Manager        []*User   `json:"managers,omitempty" bson:"managers,omitempty"`
	Type           string    `json:"type" bson:"type"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	LastActiveTime time.Time `json:"last_active_time" bson:"last_active_time"`
	Status         string    `json:"status" bson:"status"`
	LastMessage    Message   `json:"last_message" bson:"last_message"`
	IsPinned       bool      `json:"is_pinned" bson:"-"`
	Version        string    `json:"version"`
}
