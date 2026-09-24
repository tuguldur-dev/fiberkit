package auth

import "time"

type Account struct {
	ID           uint      `json:"id" gorm:"primaryKey" example:"1"`
	Email        string    `json:"email" gorm:"uniqueIndex;not null" example:"admin@example.com"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}
