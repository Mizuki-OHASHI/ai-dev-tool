package dao

import "main/model"

func GetUsers() ([]model.User, error) {
	rows, err := Db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUser(id string) (model.User, error) {
	var user model.User
	err := Db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func CreateUser(user *model.User) error {
	// ID 追加
	_, err := Db.Exec("INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())", user.ID, user.Name, user.Email, user.Password)
	return err
}

func UpdateUser(user *model.User) error {
	_, err := Db.Exec("UPDATE users SET name = ?, email = ?, password = ?, updated_at = NOW() WHERE id = ?", user.Name, user.Email, user.Password, user.ID)
	return err
}

func DeleteUser(id string) error {
	_, err := Db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
