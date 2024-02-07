package models

import (
	"context"
	"main/database"

	"github.com/jackc/pgx/v5"
)

func AlreadyExists(username string) (err error) {
	var retrievedUsername string
	query := "SELECT username FROM users WHERE username = $1;"
	err = database.Conn.QueryRow(context.Background(), query, username).Scan(&retrievedUsername)
	return
}

func InsertIntoUsers(user *User) (pgx.Rows, error) {
	query := "INSERT INTO users (username, first_name, last_name, password, token, user_type) VALUES ($1, $2, $3, $4, $5, $6);"
	row, err := database.Conn.Query(context.Background(), query, user.Username, user.FirstName, user.LastName, user.Password, user.Token, user.UserType)
	return row, err
}

func VerifyUsername(username string) (err error, foundUser User) {
	query := "SELECT * FROM users WHERE username = $1;"
	err = database.Conn.QueryRow(context.Background(), query, username).Scan(&foundUser.Username, &foundUser.FirstName, &foundUser.LastName,
		&foundUser.Password, &foundUser.Token, &foundUser.UserType)
	return
}

func UpdateToken(username string, token string) (row pgx.Rows, err error) {
	query := "UPDATE users SET token = $1 WHERE username = $2"
	row, err = database.Conn.Query(context.Background(), query, token, username)
	return
}

func RetrieveUser(username string) (user *User, err error) {
	query := "SELECT * FROM users WHERE username = $1;"
	err = database.Conn.QueryRow(context.Background(), query, username).Scan(&user.Username, &user.FirstName, &user.LastName, &user.Password, &user.Token, &user.UserType)
	return
}

func GetUserTaskList(username string) (rows pgx.Rows, err error) {
	query := "SELECT * FROM tasks WHERE belongs = $1;"
	rows, err = database.Conn.Query(context.Background(), query, username)
	return
}

func CreateUserTask(task Task) (row pgx.Rows, err error) {
	query := "INSERT INTO tasks (belongs, title, description, done) VALUES ($1, $2, $3, $4);"
	row, err = database.Conn.Query(context.Background(), query, task.Belongs, task.Title, task.Description, task.Done)
	return
}

func UpdateUserTask(task Task) (row pgx.Rows, err error) {
	query := "UPDATE tasks SET title = $1, description = $2, done = $3 WHERE taskID = $4;"
	row, err = database.Conn.Query(context.Background(), query, task.Title, task.Description, task.Done, task.TaskID)
	return
}

func DeleteUserTask(taskid string) (row pgx.Rows, err error) {
	query := "DELETE FROM tasks WHERE taskID = $1;"
	row, err = database.Conn.Query(context.Background(), query, taskid)
	return
}
