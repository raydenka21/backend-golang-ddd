package model

import (
	"gorm.io/gorm"
	"time"
)

// Movie ...
type Movie struct {
	ID        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Title     string    `json:"title"`
	Year      int       `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
