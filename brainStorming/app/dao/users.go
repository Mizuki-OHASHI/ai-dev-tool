package dao

import (
	"errors"
	"main/model"
)

func GetUsers() ([]model.User, error) {
	rows, err := Db.Query("SELECT * FROM users")
	if err != nil {
		return nil, errors.New("failed to fetch users from database")
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, errors.New("failed to scan user data")
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUser(id string) (model.User, error) {
	var user model.User
	err := Db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, errors.New("failed to fetch user from database")
	}

	return user, nil
}

func CreateUser(user *model.User) error {
	_, err := Db.Exec("INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())", user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return errors.New("failed to insert user into database")
	}
	return nil
}

func UpdateUser(user *model.User) error {
	_, err := Db.Exec("UPDATE users SET name = ?, email = ?, password = ?, updated_at = NOW() WHERE id = ?", user.Name, user.Email, user.Password, user.ID)
	if err != nil {
		return errors.New("failed to update user in database")
	}
	return nil
}

func DeleteUser(id string) error {
	_, err := Db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return errors.New("failed to delete user from database")
	}
	return nil
}

func IsUsernameExists(username string) (bool, error) {
	var exists bool
	err := Db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE name=?)", username).Scan(&exists)
	if err != nil {
		return false, errors.New("failed to check if username exists in database")
	}
	return exists, nil
}

func IsUserExists(id string) (bool, error) {
	var exists bool
	err := Db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id=?)", id).Scan(&exists)
	if err != nil {
		return false, errors.New("failed to check if user exists in database")
	}
	return exists, nil
}