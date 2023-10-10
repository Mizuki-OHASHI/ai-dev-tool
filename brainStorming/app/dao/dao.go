package dao

import (
	"database/sql"
	"fmt"
	"log"
	"main/model"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func OpenSql() {
	var (
		mysqlUser     = os.Getenv("MYSQL_USER")
		mysqlUserPwd  = os.Getenv("MYSQL_ROOT_PASSWORD")
		mysqlHost     = os.Getenv("MYSQL_HOST")
		mysqlDatabase = os.Getenv("MYSQL_DATABASE")
	)

	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", mysqlUser, mysqlUserPwd, mysqlHost, mysqlDatabase)
	Db_, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}

	if err := Db_.Ping(); err != nil {
		log.Fatalf("fail: Db_.Ping, %v\n", err)
	}

	Db = Db_
}

func CloseDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		s := <-sig
		log.Printf("received syscall, %v\n", s)

		if err := Db.Close(); err != nil {
			log.Fatal(err)
		}

		log.Printf("success: Db.Close")
		os.Exit(0)
	}()
}

func GetUsers() ([]model.User, error) {
	rows, err := Db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUser(id string) (model.User, error) {
	var user model.User
	err := Db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func CreateUser(user *model.User) error {
	// ID 追加
	_, err := Db.Exec("INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())", user.ID, user.Name, user.Email, user.Password)
	return err
}

func UpdateUser(user *model.User) error {
	_, err := Db.Exec("UPDATE users SET name = ?, email = ?, password = ?, updated_at = NOW() WHERE id = ?", user.Name, user.Email, user.Password, user.ID)
	return err
}

func DeleteUser(id string) error {
	_, err := Db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
