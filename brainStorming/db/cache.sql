CREATE TABLE users (
  id CHAR(32) PRIMARY KEY,
  name CHAR(50) NOT NULL,
  email CHAR(50) NOT NULL,
  password CHAR(50) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL
);

INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES 
("12345678901234567890123456789012", 'John Doe', 'johndoe@example.com', 'password123', '2022-01-01 00:00:00', '2022-01-01 00:00:00');

CREATE USER 'app'@'%' identified by 'password';

GRANT all ON brain_storming.* TO 'app'@'%';