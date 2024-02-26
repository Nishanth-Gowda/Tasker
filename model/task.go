package model

import (

	"gorm.io/gorm"
)

type Task struct {
    gorm.Model
    Name        string `json:"name" gorm:"not null"`
    Status      string `json:"status" gorm:"not null"`
    ProjectID   uint   `json:"projectID" gorm:"not null"`
    AssignedToID uint  `json:"assignedTo" gorm:"not null"`
}