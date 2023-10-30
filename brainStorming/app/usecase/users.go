package usecase

import (
	"errors"
	"main/dao"
	"main/model"

	"github.com/google/uuid"
)

func GetUsers() ([]model.User, error) {
	users, err := dao.GetUsers()
	if err != nil {
		return nil, errors.New("failed to get users")
	}
	return users, nil
}

func GetUser(id string) (model.User, error) {
	user, err := dao.GetUser(id)
	if err != nil {
		return model.User{}, errors.New("failed to get user")
	}
	return user, nil
}

func CreateUser(user *model.User) error {
	exists, err := dao.IsUsernameExists(user.Name)
	if err != nil {
		return errors.New("failed to check if username exists")
	}
	if exists {
		return errors.New("username already exists")
	}

	user.ID = uuid.New().String()
	err = dao.CreateUser(user)
	if err != nil {
		return errors.New("failed to create user")
	}
	return nil
}

func UpdateUser(user *model.User) error {
	err := dao.UpdateUser(user)
	if err != nil {
		return errors.New("failed to update user")
	}

	return nil
}

func DeleteUser(id string) error {
	err := dao.DeleteUser(id)
	if err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}
