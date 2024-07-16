package routes

import (
	"todoapp/controllers"
	"todoapp/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	todoRoutes := r.Group("/todos")
	todoRoutes.Use(middleware.AuthMiddleware())
	{
		todoRoutes.GET("/", controllers.GetToDos)
		todoRoutes.POST("/", controllers.CreateToDo)
		todoRoutes.PUT("/:id", controllers.UpdateToDo)
		todoRoutes.DELETE("/:id", controllers.DeleteToDo)
	}
}
