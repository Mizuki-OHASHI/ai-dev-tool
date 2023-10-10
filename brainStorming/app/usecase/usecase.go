package usecase

import (
	"main/dao"
	"main/model"
)

func GetUser(id int) (*model.User, error) {
	return dao.GetUser(id)
}