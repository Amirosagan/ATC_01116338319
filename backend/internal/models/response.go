package models

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

// SuccessResponse represents a success response with a message
type SuccessResponse struct {
	Message string `json:"message" example:"operation successful"`
} 