package initializers

import (
	"errors"
	"log"
	"strings"

	"IkonKutz.API/models"
	"IkonKutz.API/utils"
	"gorm.io/gorm"
)

func SeedAdminUser() {
	email := strings.TrimSpace(strings.ToLower(AppConfig.AdminEmail))
	password := strings.TrimSpace(AppConfig.AdminPassword)
	name := strings.TrimSpace(AppConfig.AdminName)

	// If admin env vars are not provided, just skip seeding.
	if email == "" || password == "" {
		log.Println("ADMIN_EMAIL or ADMIN_PASSWORD not set, skipping admin seeding")
		return
	}

	if name == "" {
		name = "Victor" // Default admin name if not provided
	}

	var existingUser models.User
	err := DB.Where("email = ?", email).First(&existingUser).Error

	if err == nil {
		// User already exists.
		if existingUser.Role != "admin" {
			existingUser.Role = "admin"

			if saveErr := DB.Save(&existingUser).Error; saveErr != nil {
				log.Println("failed to upgrade existing user to admin:", saveErr)
				return
			}

			log.Println("existing user upgraded to admin successfully")
			return
		}

		log.Println("admin user already exists")
		return
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("failed to check for existing admin user:", err)
		return
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Println("failed to hash admin password:", err)
		return
	}

	admin := models.User{
		Name:         name,
		Email:        email,
		PasswordHash: hashedPassword,
		Role:         "admin",
	}

	if err := DB.Create(&admin).Error; err != nil {
		log.Println("failed to create admin user:", err)
		return
	}

	log.Println("admin user seeded successfully")
}
