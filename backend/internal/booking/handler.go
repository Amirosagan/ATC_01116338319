package booking

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"online-task/internal/models"
	"online-task/pkg/database"
)

// @Summary Create a booking
// @Description Book an event for the authenticated user
// @Tags bookings
// @Accept json
// @Produce json
// @Param request body models.CreateBookingRequest true "Booking details"
// @Security Bearer
// @Success 201 {object} models.BookingResponse
// @Router /bookings [post]
func CreateBookingHandler(c *gin.Context) {
	var req models.CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")

	// Check if event exists
	var event models.Event
	if err := database.GetDB().First(&event, "id = ?", req.EventID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch event"})
		return
	}

	// Check if user already has a booking for this event
	var existingBooking models.Booking
	err := database.GetDB().Where("user_id = ? AND event_id = ?", userID, req.EventID).First(&existingBooking).Error
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "You have already booked this event"})
		return
	} else if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing booking"})
		return
	}

	booking := models.Booking{
		ID:      uuid.New().String(),
		UserID:  userID.(string),
		EventID: req.EventID,
	}

	if err := database.GetDB().Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	// Load the event details for the response
	if err := database.GetDB().Preload("Event").First(&booking, "id = ?", booking.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch booking details"})
		return
	}

	response := models.BookingResponse{
		ID:        booking.ID,
		UserID:    booking.UserID,
		EventID:   booking.EventID,
		Event:     booking.Event,
		CreatedAt: booking.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// @Summary Get user bookings
// @Description Get all bookings for the authenticated user
// @Tags bookings
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {array} models.BookingResponse
// @Router /bookings/user [get]
func GetUserBookingsHandler(c *gin.Context) {
	userID, _ := c.Get("userID")

	var bookings []models.Booking
	if err := database.GetDB().Preload("Event.Tags").Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	var response []models.BookingResponse
	for _, booking := range bookings {
		response = append(response, models.BookingResponse{
			ID:        booking.ID,
			UserID:    booking.UserID,
			EventID:   booking.EventID,
			Event:     booking.Event,
			CreatedAt: booking.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
} 