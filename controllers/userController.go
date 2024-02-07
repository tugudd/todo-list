package controllers

import (
	"context"
	"fmt"
	"main/database"
	"main/helpers"
	"main/models"
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn = database.DBinstance()

func Signup(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.AlreadyExists(user.Username)

	if err == nil {
		c.JSON(500, gin.H{"error": "User with this username already exists"})
		return
	}

	pass := service.HashPassword(user.Password)
	user.Password = pass

	token, err := helpers.GenerateAllTokens(user.Username, user.FirstName, user.LastName, user.UserType)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not create token"})
		return
	}
	user.Token = token

	row, err := models.InsertIntoUsers(&user)
	row.Close()

	if err != nil {
		c.JSON(500, gin.H{"error": "User was not created"})
		return
	}
	c.JSON(200, "User was successfully created")
}

func Login(c *gin.Context) {
	var user models.User
	var foundUser models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err, foundUser := models.VerifyUsername(user.Username)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"error": "Wrong username or password"})
		return
	}

	validPassword := service.VerifyPassword(user.Password, foundUser.Password)
	if validPassword == false {
		c.JSON(500, gin.H{"error": "Wrong username or password"})
		return
	}

	token, err := helpers.GenerateAllTokens(foundUser.FirstName, foundUser.FirstName, foundUser.LastName, foundUser.UserType)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not create token"})
		return
	}

	foundUser.Token = token
	row, err := models.UpdateToken(foundUser.Username, token)
	row.Close()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, foundUser)
}

func GetUsers(c *gin.Context) {
	if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var users []models.User

	query := "SELECT * FROM users;"
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Username, &user.FirstName, &user.LastName, &user.Password, &user.Token, &user.UserType)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		users = append(users, user)
	}

	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	username := c.Param("username")
	if err := helpers.MatchUserTypeToUsername(c, username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := models.RetrieveUser(username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func GetList(c *gin.Context) {
	username := c.Param("username")
	if err := helpers.MatchUserTypeToUsername(c, username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rows, err := models.GetUserTaskList(username)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasklist []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.TaskID, &task.Belongs, &task.Title, &task.Description, &task.Done)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		tasklist = append(tasklist, task)
	}

	c.JSON(200, tasklist)
}

func CreateTask(c *gin.Context) {
	username := c.Param("username")
	if err := helpers.MatchUserTypeToUsername(c, username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(500, gin.H{"errorr": err.Error()})
		return
	}

	msg := models.ValidateTask(task, false)
	if msg != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	row, err := models.CreateUserTask(task)
	row.Close()

	if err != nil {
		c.JSON(500, gin.H{"errorr": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "Task was successfuly created"})
}

func UpdateTask(c *gin.Context) {
	username := c.Param("username")
	taskid := c.Param("task_id")

	if err := helpers.MatchUserTypeToUsername(c, username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	task.TaskID = taskid
	task.Belongs = username

	msg := models.ValidateTask(task, true)
	if msg != "" {
		c.JSON(400, gin.H{"error": msg})
		return
	}

	row, err := models.UpdateUserTask(task)
	row.Close()

	if err != nil {
		c.JSON(500, gin.H{"errorr": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": task})

}

func DeleteTask(c *gin.Context) {
	username := c.Param("username")
	taskid := c.Param("task_id")

	if err := helpers.MatchUserTypeToUsername(c, username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row, err := models.DeleteUserTask(taskid)
	row.Close()

	if err != nil {
		c.JSON(500, gin.H{"errorr": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "Task was successfuly deleted"})
}
