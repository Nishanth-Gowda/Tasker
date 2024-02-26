package model

import (
	"time"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ID        uint     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null; column:name; size:255"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
}
