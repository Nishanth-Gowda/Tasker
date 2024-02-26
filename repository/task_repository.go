package repository

import (
	"github.com/nishanth-gowda/tasker/model"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		DB: db,
	}
}

func (tr *TaskRepository) CreateTask(task *model.Task) error {
	result := tr.DB.Create(task)
	return result.Error
}

func (tr *TaskRepository) GetTasksByID(id uint) (model.Task, error) {
	var task model.Task
	err := tr.DB.First(&task, id).Error
	return task, err
}

func (tr *TaskRepository) GetTasksByProject(projectID uint) ([]model.Task, error) {
	var tasks []model.Task
	err := tr.DB.Where("project_id = ?", projectID).Find(&tasks).Error
	return tasks, err
}
