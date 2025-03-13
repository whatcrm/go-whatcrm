package models

type User struct {
	ID     string `json:"id"`
	ChatID string `json:"chat_id"`
	Domain string `json:"domain"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	Status string `json:"status"`
	Email  string `json:"email"`
}
