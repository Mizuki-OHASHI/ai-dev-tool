package usecase

import (
	"main/dao"
	"main/model"
)

func GetPosts() ([]model.Post, error) {
	return dao.GetPosts()
}

func GetPost(id string) (model.Post, error) {
	return dao.GetPost(id)
}

func CreatePost(post *model.Post) error {
	return dao.CreatePost(post)
}

func UpdatePost(post *model.Post) error {
	return dao.UpdatePost(post)
}

func DeletePost(id string) error {
	return dao.DeletePost(id)
}

func GetPostsByUser(userId string) ([]model.Post, error) {
	return dao.GetPostsByUser(userId)
}
