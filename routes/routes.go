package routes

import (
	"todoapp/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/todos", controllers.GetToDos)
	router.POST("/todos", controllers.CreateToDo)
	router.PUT("/todos/:id", controllers.UpdateToDo)
	router.DELETE("/todos/:id", controllers.DeleteToDo)
}
