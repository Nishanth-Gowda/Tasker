package repository

import (
	"github.com/nishanth-gowda/tasker/model"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	DB *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		DB: db,
	}
}

func (r *ProjectRepository) CreateProject(project *model.Project) error {
	result := r.DB.Create(project)
	return result.Error
}

func (r *ProjectRepository) GetProject(id uint) (*model.Project, error) {
	var project model.Project
	err := r.DB.First(&project, id).Error
	return &project, err
}

func (r *ProjectRepository) DeleteProject(id uint) error {
	result := r.DB.Delete(&model.Project{}, id)
	return result.Error
}