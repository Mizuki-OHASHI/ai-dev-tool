package model

/*
CREATE TABLE users (
  id CHAR(32) PRIMARY KEY,
  name CHAR(50) NOT NULL,
  email CHAR(50) NOT NULL,
  password CHAR(50) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL
);
*/

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
