package dao

import "main/model"

func GetPosts() ([]model.Post, error) {
	rows, err := Db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.PostedBy, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetPost(id string) (model.Post, error) {
	var post model.Post
	err := Db.QueryRow("SELECT * FROM posts WHERE id = ?", id).Scan(&post.ID, &post.PostedBy, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return model.Post{}, err
	}

	return post, nil
}

func CreatePost(post *model.Post) error {
	_, err := Db.Exec("INSERT INTO posts (id, posted_by, title, body, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())", post.ID, post.PostedBy, post.Title, post.Body)
	return err
}

func UpdatePost(post *model.Post) error {
	_, err := Db.Exec("UPDATE posts SET posted_by = ?, title = ?, body = ?, updated_at = NOW() WHERE id = ?", post.PostedBy, post.Title, post.Body, post.ID)
	return err
}

func DeletePost(id string) error {
	_, err := Db.Exec("DELETE FROM posts WHERE id = ?", id)
	return err
}

func GetPostsByUser(userId string) ([]model.Post, error) {
	rows, err := Db.Query("SELECT * FROM posts WHERE posted_by = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.PostedBy, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}