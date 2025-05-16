package entities

import "time"

type Product struct {
	ID        int `gorm:"primary_key"`
	Name      string
	Price     float64
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
