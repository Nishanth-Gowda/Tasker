package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nishanth-gowda/tasker/handlers"
)

func RegisterRoutes(r *gin.Engine, projectHandler *handlers.ProjectHandler, taskHandler *handlers.TaskHandler) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/project/:id", projectHandler.GetProject)
		v1.POST("/projects", projectHandler.CreateProject)
		v1.DELETE("/project/:id", projectHandler.DeleteProject)

		v1.GET("/tasks/:id", taskHandler.GetTasksByID)
		v1.POST("/tasks", taskHandler.CreateTask)
		v1.GET("/projects/:projectID/tasks", taskHandler.GetTasksByProject)
	}
}