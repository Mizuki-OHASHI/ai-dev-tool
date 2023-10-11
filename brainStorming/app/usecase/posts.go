package usecase

import (
	"errors"
	"main/dao"
	"main/model"
)

func GetPosts() ([]model.Post, error) {
	posts, err := dao.GetPosts()
	if err != nil {
		return nil, errors.New("failed to get posts")
	}
	return posts, nil
}

func GetPost(id string) (model.Post, error) {
	post, err := dao.GetPost(id)
	if err != nil {
		return model.Post{}, errors.New("failed to get post")
	}
	return post, nil
}

func CreatePost(post *model.Post) error {
	err := dao.CreatePost(post)
	if err != nil {
		return errors.New("failed to create post")
	}
	return nil
}

func UpdatePost(post *model.Post) error {
	err := dao.UpdatePost(post)
	if err != nil {
		return errors.New("failed to update post")
	}
	return nil
}

func DeletePost(id string) error {
	err := dao.DeletePost(id)
	if err != nil {
		return errors.New("failed to delete post")
	}
	return nil
}

func GetPostsByUser(userId string) ([]model.Post, error) {
	posts, err := dao.GetPostsByUser(userId)
	if err != nil {
		return nil, errors.New("failed to get posts by user")
	}
	return posts, nil
}
