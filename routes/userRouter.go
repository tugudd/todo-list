package routes

import (
	controller "main/controllers"
	middleware "main/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers)
	incomingRoutes.GET("/users/:username", controller.GetUser)
	incomingRoutes.GET("/users/:username/tasklist", controller.GetList)
	incomingRoutes.POST("/users/:username/tasklist", controller.CreateTask)
	incomingRoutes.PUT("/users/:username/tasklist/:task_id", controller.UpdateTask)
	incomingRoutes.DELETE("/users/:username/tasklist/:task_id", controller.DeleteTask)
}
