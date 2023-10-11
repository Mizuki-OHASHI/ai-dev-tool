package usecase

import (
	"main/dao"
	"main/model"
	"errors"
)

func GetUsers() ([]model.User, error) {
	users, err := dao.GetUsers()
	if err != nil {
		return nil, errors.New("Failed to get users")
	}
	return users, nil
}

func GetUser(id string) (model.User, error) {
	user, err := dao.GetUser(id)
	if err != nil {
		return model.User{}, errors.New("Failed to get user")
	}
	return user, nil
}

func CreateUser(user *model.User) error {
	err := dao.CreateUser(user)
	if err != nil {
		return errors.New("Failed to create user")
	}
	return nil
}

func UpdateUser(user *model.User) error {
	err := dao.UpdateUser(user)
	if err != nil {
		return errors.New("Failed to update user")
	}
	return nil
}

func DeleteUser(id string) error {
	err := dao.DeleteUser(id)
	if err != nil {
		return errors.New("Failed to delete user")
	}
	return nil
}