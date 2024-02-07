package main

import (
	"context"
	"fmt"
	"log"
	"main/database"
	"main/helpers"
	"main/models"
	"main/routes"
	"main/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	conn := database.Conn
	row, err := conn.Query(context.Background(), models.UserTableQuery)
	row.Close()
	defer conn.Close(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	row, err = conn.Query(context.Background(), models.TaskTableQuery)
	row.Close()

	if err != nil {
		log.Fatal(err)
	}

	var username string
	query := "SELECT username FROM users WHERE username = $1"
	err = conn.QueryRow(context.Background(), query, "admin").Scan(&username)

	if err != nil {
		fmt.Println(err.Error())
		token, _ := helpers.GenerateAllTokens("admin", "Yslam", "Orazov", "ADMIN")
		pass := service.HashPassword("secret")

		query = "INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6);"
		row, err = conn.Query(context.Background(), query, "admin", "Yslam", "Orazov", pass, token, "ADMIN")
		row.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted to test page"})
	})

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Run(":8081")
}
