package models

type User struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	UserType  string `json:"user_type"`
}

var UserTableQuery = `CREATE TABLE IF NOT EXISTS users (
	username VARCHAR(100) PRIMARY KEY NOT NULL,
	first_name VARCHAR(100),
	last_name VARCHAR(100),
	password VARCHAR(75) NOT NULL,
	token VARCHAR(1000) NOT NULL,
	user_type VARCHAR(20) NOT NULL
);`
