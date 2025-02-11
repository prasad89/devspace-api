package models

import (
	"github.com/prasad89/devspace-api/initializers"
)

// User struct to hold user details
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

// GetByUsername fetches a user by username
func GetByUsername(username string) (User, error) {
	var user User
	result := initializers.DB.Where("username = ?", username).First(&user)
	return user, result.Error
}
