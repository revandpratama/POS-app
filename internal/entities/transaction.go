package entities

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        int `gorm:"primary_key"`
	ProductID int
	UserID    int
	Quantity  int
	Total     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
