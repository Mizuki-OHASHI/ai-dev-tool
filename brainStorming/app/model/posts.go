package model

/*
CREATE TABLE posts (
  id CHAR(32) PRIMARY KEY,
  posted_by CHAR(32) NOT NULL,
  title CHAR(50) NOT NULL,
  body CHAR(140) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL
);
*/

type Post struct {
	ID        string `json:"id"`
	PostedBy  string `json:"posted_by"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
