package routes

import (
	"TaskManager/handlers"
	"TaskManager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.RegisterUser)
		auth.POST("/login", handlers.LoginUser)
	}

	tasks := api.Group("/tasks")
	tasks.Use(middleware.RequireAuth)
	{
		tasks.GET("", handlers.GetTasks)
		tasks.POST("", handlers.CreateTask)
		tasks.GET("/:id", handlers.GetTask)
		tasks.PUT("/:id", handlers.UpdateTask)
		tasks.DELETE("/:id", handlers.DeleteTask)
		//tasks.GET("", handlers.FilterByStatus)
		tasks.PATCH("/:id", handlers.PatchTask)
	}
}
