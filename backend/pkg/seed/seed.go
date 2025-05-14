package seed

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"online-task/internal/models"
	"online-task/pkg/database"
)

const (
	AdminEmail    = "admin@example.com"
	AdminPassword = "admin123"
	AdminUsername = "admin"
)

func SeedAdminUser() error {
	db := database.GetDB()

	// Check if admin already exists
	var existingUser models.User
	if err := db.Where("email = ?", AdminEmail).First(&existingUser).Error; err == nil {
		log.Println("Admin user already exists")
		return nil
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(AdminPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create admin user
	adminUser := models.User{
		Email:    AdminEmail,
		Username: AdminUsername,
		Password: string(hashedPassword),
		Role:     "admin",
	}

	if err := db.Create(&adminUser).Error; err != nil {
		return err
	}

	log.Printf("Admin user created successfully with email: %s and password: %s\n", AdminEmail, AdminPassword)
	return nil
} 