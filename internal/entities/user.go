package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int `gorm:"primary_key"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
