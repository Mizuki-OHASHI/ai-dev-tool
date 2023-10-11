package model

type Post struct {
	ID        string `json:"id"`
	PostedBy  string `json:"posted_by"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}