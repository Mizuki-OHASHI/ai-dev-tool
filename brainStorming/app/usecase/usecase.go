package usecase

import (
	"main/dao"
	"main/model"
)

func GetUsers() ([]model.User, error) {
	return dao.GetUsers()
}

func GetUser(id string) (model.User, error) {
	return dao.GetUser(id)
}

func CreateUser(user *model.User) error {
	return dao.CreateUser(user)
}

func UpdateUser(user *model.User) error {
	return dao.UpdateUser(user)
}

func DeleteUser(id string) error {
	return dao.DeleteUser(id)
}
