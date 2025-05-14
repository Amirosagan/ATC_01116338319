package upload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	maxFileSize    = 3 * 1024 * 1024 // 3MB
	uploadDir      = "./uploads/images"
	allowedFormats = ".jpg,.jpeg,.png,.gif"
)

type UploadResponse struct {
	ImageURL string `json:"imageUrl"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func validateFileType(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return strings.Contains(allowedFormats, ext)
}

// @Summary Upload an image
// @Description Upload an image file (max 3MB, formats: jpg, jpeg, png, gif)
// @Tags upload
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Image file"
// @Security Bearer
// @Success 200 {object} UploadResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /upload/image [post]
func UploadImageHandler(c *gin.Context) {
	// Create uploads directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}
	defer file.Close()

	// Check file size
	if header.Size > maxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds 3MB limit"})
		return
	}

	// Validate file type
	if !validateFileType(header.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file format. Allowed formats: jpg, jpeg, png, gif"})
		return
	}

	// Generate unique filename
	ext := filepath.Ext(header.Filename)
	newFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filepath := filepath.Join(uploadDir, newFilename)

	// Create new file
	dst, err := os.Create(filepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file"})
		return
	}
	defer dst.Close()

	// Copy file content
	if _, err := io.Copy(dst, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Return the URL
	imageURL := fmt.Sprintf("/uploads/images/%s", newFilename)
	c.JSON(http.StatusOK, UploadResponse{ImageURL: imageURL})
}
