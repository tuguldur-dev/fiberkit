package user

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	Name      string    `json:"name" gorm:"not null" example:"Ada Lovelace"`
	CreatedAt time.Time `json:"created_at"`
}
