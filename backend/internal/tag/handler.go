package tag

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"online-task/internal/models"
	"online-task/pkg/database"
)

// @Summary Get all tags
// @Description Get a list of all tags
// @Tags tags
// @Accept json
// @Produce json
// @Success 200 {array} models.Tag
// @Router /tags [get]
func GetAllTagsHandler(c *gin.Context) {
	var tags []models.Tag
	if err := database.GetDB().Find(&tags).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tags"})
		return
	}

	c.JSON(http.StatusOK, tags)
}

// @Summary Get tag by ID
// @Description Get details of a specific tag
// @Tags tags
// @Accept json
// @Produce json
// @Param id path string true "Tag ID"
// @Success 200 {object} models.Tag
// @Router /tags/{id} [get]
func GetTagHandler(c *gin.Context) {
	id := c.Param("id")
	var tag models.Tag

	if err := database.GetDB().First(&tag, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tag"})
		return
	}

	c.JSON(http.StatusOK, tag)
}

// @Summary Create a new tag
// @Description Create a new tag (admin only)
// @Tags tags
// @Accept json
// @Produce json
// @Param request body models.CreateTagRequest true "Tag details"
// @Security Bearer
// @Success 201 {object} models.Tag
// @Router /tags [post]
func CreateTagHandler(c *gin.Context) {
	var req models.CreateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag := models.Tag{
		ID:   uuid.New().String(),
		Name: req.Name,
	}

	if err := database.GetDB().Create(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tag"})
		return
	}

	c.JSON(http.StatusCreated, tag)
}

// @Summary Update a tag
// @Description Update an existing tag (admin only)
// @Tags tags
// @Accept json
// @Produce json
// @Param id path string true "Tag ID"
// @Param request body models.CreateTagRequest true "Tag details"
// @Security Bearer
// @Success 200 {object} models.Tag
// @Router /tags/{id} [put]
func UpdateTagHandler(c *gin.Context) {
	id := c.Param("id")
	var req models.CreateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tag models.Tag
	if err := database.GetDB().First(&tag, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tag"})
		return
	}

	tag.Name = req.Name

	if err := database.GetDB().Save(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tag"})
		return
	}

	c.JSON(http.StatusOK, tag)
}

// @Summary Delete a tag
// @Description Delete an existing tag (admin only)
// @Tags tags
// @Accept json
// @Produce json
// @Param id path string true "Tag ID"
// @Security Bearer
// @Success 200 {object} models.SuccessResponse
// @Router /tags/{id} [delete]
func DeleteTagHandler(c *gin.Context) {
	id := c.Param("id")

	// Start transaction
	tx := database.GetDB().Begin()

	// Clear event associations
	if err := tx.Table("event_tags").Where("tag_id = ?", id).Delete(&struct{}{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear tag associations"})
		return
	}

	// Delete the tag
	if err := tx.Where("id = ?", id).Delete(&models.Tag{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tag"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Tag deleted successfully"})
} 