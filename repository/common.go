package repository

import "github.com/nishanth-gowda/tasker/model"

type Common interface {
	//project
	GetProject(id int) (*model.Project, error)
	CreateProject(project *model.Project) error
	DeleteProject(id int) error

	//tasks
	GetTasksByID(projectID int) ([]model.Task, error)
	CreateTask(task *model.Task) error
	GetTasksByProject(projectID int) ([]model.Task, error)
}
