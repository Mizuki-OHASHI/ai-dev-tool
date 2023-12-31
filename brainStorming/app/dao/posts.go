package dao

import (
	"errors"
	"main/model"
)

func GetPosts() ([]model.Post, error) {
	rows, err := Db.Query("SELECT posts.*, users.name FROM posts JOIN users ON posts.posted_by = users.id")
	if err != nil {
		return nil, errors.New("failed to fetch posts from database")
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.PostedBy, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt, &post.UserName); err != nil {
			return nil, errors.New("failed to scan post data")
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetPost(id string) (model.Post, error) {
	var post model.Post
	err := Db.QueryRow("SELECT posts.*, users.name FROM posts JOIN users ON posts.posted_by = users.id WHERE posts.id = ?", id).Scan(&post.ID, &post.PostedBy, &post.UserName, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return model.Post{}, errors.New("failed to fetch post from database")
	}

	return post, nil
}

func CreatePost(post *model.Post) error {
	_, err := Db.Exec("INSERT INTO posts (id, posted_by, title, body, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())", post.ID, post.PostedBy, post.Title, post.Body)
	if err != nil {
		return errors.New("failed to insert post into database")
	}
	return nil
}

func UpdatePost(post *model.Post) error {
	_, err := Db.Exec("UPDATE posts SET posted_by = ?, title = ?, body = ?, updated_at = NOW() WHERE id = ?", post.PostedBy, post.Title, post.Body, post.ID)
	if err != nil {
		return errors.New("failed to update post in database")
	}
	return nil
}

func DeletePost(id string) error {
	_, err := Db.Exec("DELETE FROM posts WHERE id = ?", id)
	if err != nil {
		return errors.New("failed to delete post from database")
	}
	return nil
}

func GetPostsByUser(userId string) ([]model.Post, error) {
	rows, err := Db.Query("SELECT posts.*, users.name FROM posts JOIN users ON posts.posted_by = users.id WHERE posts.posted_by = ?", userId)
	if err != nil {
		return nil, errors.New("failed to fetch posts by user from database")
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.PostedBy, &post.UserName, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, errors.New("failed to scan post data")
		}
		posts = append(posts, post)
	}

	return posts, nil
}
