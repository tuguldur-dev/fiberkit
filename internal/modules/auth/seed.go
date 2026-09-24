package auth

import (
	"log"

	"fiberkit/internal/database"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	var count int64
	if err := database.DB.Model(&Account{}).Count(&count).Error; err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	if err := database.DB.Create(&Account{
		Email:        "admin@example.com",
		PasswordHash: string(hash),
	}).Error; err != nil {
		log.Fatal(err)
	}

	log.Println("seeded admin@example.com / admin")
}
