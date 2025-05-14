package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID          string         `gorm:"primarykey" json:"id"`
	UserID      string         `gorm:"not null" json:"userId"`
	EventID     string         `gorm:"not null" json:"eventId"`
	BookingDate time.Time      `json:"bookingDate" format:"date-time" example:"2024-03-20T10:00:00Z"`
	CreatedAt   time.Time      `json:"createdAt" format:"date-time" example:"2024-03-20T10:00:00Z"`
	UpdatedAt   time.Time      `json:"updatedAt" format:"date-time" example:"2024-03-20T10:00:00Z"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Event       Event          `json:"event,omitempty"`
	User        User           `json:"user,omitempty"`
}

type CreateBookingRequest struct {
	EventID string `json:"eventId" binding:"required"`
}

type BookingResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	EventID   string    `json:"eventId"`
	Event     Event     `json:"event"`
	CreatedAt time.Time `json:"createdAt"`
} 