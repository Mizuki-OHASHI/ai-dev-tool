CREATE TABLE users (
  id CHAR(36) PRIMARY KEY,
  name CHAR(50) NOT NULL,
  email CHAR(50) NOT NULL,
  password CHAR(50) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL
);

CREATE TABLE posts (
  id CHAR(36) PRIMARY KEY,
  posted_by CHAR(36) NOT NULL,
  title CHAR(50) NOT NULL,
  body CHAR(140) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL
);
