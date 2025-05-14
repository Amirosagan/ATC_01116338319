package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          string         `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Date        time.Time      `json:"date" format:"date-time" example:"2024-03-20T15:00:00Z"`
	Location    string         `json:"location"`
	Price       float64        `json:"price"`
	Image       string         `json:"image"`
	CreatedAt   time.Time      `json:"createdAt" format:"date-time" example:"2024-03-20T10:00:00Z"`
	UpdatedAt   time.Time      `json:"updatedAt" format:"date-time" example:"2024-03-20T10:00:00Z"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Tags        []Tag          `gorm:"many2many:event_tags;" json:"tags,omitempty"`
}

type Tag struct {
	ID        string         `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"unique;not null" json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type CreateEventRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	Date        time.Time `json:"date" binding:"required" format:"date-time" example:"2024-03-20T15:00:00Z"`
	Location    string    `json:"location" binding:"required"`
	Price       float64   `json:"price" binding:"required,min=0"`
	Image       string    `json:"image" binding:"required"`
	TagIDs      []string  `json:"tagIds,omitempty"`
} 