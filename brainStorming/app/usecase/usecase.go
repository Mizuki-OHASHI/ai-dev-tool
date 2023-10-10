package usecase

import (
	"main/dao"
	"main/model"
)

func GetUser(id string) (*model.User, error) {
	return dao.GetUser(id)
}
