package dao

import (
	"main/model"
)

func GetUser(id int) (*model.User, error) {
	row := Db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}