package entities

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        int `gorm:"primary_key"`
	Name      string
	Price     float64
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
