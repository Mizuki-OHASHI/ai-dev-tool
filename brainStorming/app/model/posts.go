package model

import "github.com/google/uuid"

type Post struct {
	ID        string `json:"id"`
	PostedBy  string `json:"posted_by"`
	UserName  string `json:"user_name"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GeneratePostID() string {
	return uuid.New().String()
}
