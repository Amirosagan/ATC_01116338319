package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"online-task/internal/models"
	"online-task/pkg/database"
)

// @Summary Get all events
// @Description Get a list of all events
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {array} models.Event
// @Router /events [get]
func GetAllEventsHandler(c *gin.Context) {
	var events []models.Event
	if err := database.GetDB().Preload("Tags").Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch events"})
		return
	}

	c.JSON(http.StatusOK, events)
}

// @Summary Get event by ID
// @Description Get details of a specific event
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Success 200 {object} models.Event
// @Router /events/{id} [get]
func GetEventHandler(c *gin.Context) {
	id := c.Param("id")
	var event models.Event

	if err := database.GetDB().Preload("Tags").First(&event, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch event"})
		return
	}

	c.JSON(http.StatusOK, event)
}

// @Summary Create a new event
// @Description Create a new event (admin only)
// @Tags events
// @Accept json
// @Produce json
// @Param request body models.CreateEventRequest true "Event details"
// @Security Bearer
// @Success 201 {object} models.Event
// @Router /events [post]
func CreateEventHandler(c *gin.Context) {
	var req models.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event := models.Event{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		Date:        req.Date,
		Location:    req.Location,
		Price:       req.Price,
		Image:       req.Image,
	}

	// Start transaction
	tx := database.GetDB().Begin()

	if err := tx.Create(&event).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	// Add tags if provided
	if len(req.TagIDs) > 0 {
		var tags []models.Tag
		if err := tx.Find(&tags, "id IN ?", req.TagIDs).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tags"})
			return
		}
		if err := tx.Model(&event).Association("Tags").Append(tags); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add tags"})
			return
		}
	}

	tx.Commit()
	c.JSON(http.StatusCreated, event)
}

// @Summary Update an event
// @Description Update an existing event (admin only)
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Param request body models.CreateEventRequest true "Event details"
// @Security Bearer
// @Success 200 {object} models.Event
// @Router /events/{id} [put]
func UpdateEventHandler(c *gin.Context) {
	id := c.Param("id")
	var req models.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var event models.Event
	if err := database.GetDB().First(&event, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch event"})
		return
	}

	// Update fields
	event.Name = req.Name
	event.Description = req.Description
	event.Category = req.Category
	event.Date = req.Date
	event.Location = req.Location
	event.Price = req.Price
	event.Image = req.Image

	// Start transaction
	tx := database.GetDB().Begin()

	if err := tx.Save(&event).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}

	// Update tags if provided
	if len(req.TagIDs) > 0 {
		// Clear existing tags
		if err := tx.Model(&event).Association("Tags").Clear(); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear tags"})
			return
		}

		// Add new tags
		var tags []models.Tag
		if err := tx.Find(&tags, "id IN ?", req.TagIDs).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tags"})
			return
		}
		if err := tx.Model(&event).Association("Tags").Append(tags); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add tags"})
			return
		}
	}

	tx.Commit()
	c.JSON(http.StatusOK, event)
}

// @Summary Delete an event
// @Description Delete an existing event (admin only)
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Security Bearer
// @Router /events/{id} [delete]
func DeleteEventHandler(c *gin.Context) {
	id := c.Param("id")

	// Start transaction
	tx := database.GetDB().Begin()

	// Delete associated bookings first
	if err := tx.Where("event_id = ?", id).Delete(&models.Booking{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete associated bookings"})
		return
	}

	// Clear tag associations
	var event models.Event
	if err := tx.First(&event, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch event"})
		return
	}

	if err := tx.Model(&event).Association("Tags").Clear(); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear tags"})
		return
	}

	// Delete the event
	if err := tx.Delete(&event).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
} 