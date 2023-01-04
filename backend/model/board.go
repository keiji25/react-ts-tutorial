package model

import "time"

type Board struct {
	ID        uint64 `gorm:"primary_key"`
	Title     string `json:"title" gorm:"type:varchar(100);not null"`
	Content   string `json:"content" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
