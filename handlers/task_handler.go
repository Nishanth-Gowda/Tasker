package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nishanth-gowda/tasker/model"
	"github.com/nishanth-gowda/tasker/repository"
)

type TaskHandler struct {
	repository *repository.TaskRepository
}

func NewTaskHandler(repo *repository.TaskRepository) *TaskHandler {
	return &TaskHandler{
		repository: repo,
	}
}

func (th *TaskHandler) CreateTask(c *gin.Context) {
	var task model.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateTaskPayload(&task); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if err := th.repository.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

func (th *TaskHandler) GetTasksByID(c *gin.Context) {
	taskid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Task ID"})
		return
	}

	task, err := th.repository.GetTasksByID(uint(taskid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (th *TaskHandler) GetTasksByProject(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("projectID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	tasks, err := th.repository.GetTasksByProject(uint(projectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func validateTaskPayload(task *model.Task) error {
	if task.Name == "" {
		return errors.New("task name required")
	}

	if task.ProjectID == 0 {
		return errors.New("error Project ID required")
	}

	if task.AssignedToID == 0 {
		return errors.New("error Assignee id required")
	}

	return nil
}