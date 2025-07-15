package routes

import (
	"todo-app/backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(r *gin.Engine, controller *controllers.TaskController) {
	tasks := r.Group("/api/tasks")
	{
		tasks.POST("", controller.CreateTask)
		tasks.GET("", controller.GetTasks)
		tasks.PUT("/:id", controller.UpdateTask)
		tasks.DELETE("/:id", controller.DeleteTask)
	}
}
