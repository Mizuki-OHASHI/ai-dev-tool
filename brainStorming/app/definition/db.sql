CREATE TABLE users (
  id CHAR(32) PRIMARY KEY,
  name CHAR(50) NOT NULL,
  email CHAR(50) NOT NULL,
  password CHAR(50) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL
);