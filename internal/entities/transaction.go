package entities

import "time"

type Transaction struct {
	ID        int `gorm:"primary_key"`
	ProductID int
	UserID    int
	Quantity  int
	Total     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
