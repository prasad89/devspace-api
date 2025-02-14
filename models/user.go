package models

import "gorm.io/gorm"

// User struct to hold user details
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

// GetByUsername fetches a user by username using a passed database instance
func GetByUsername(db *gorm.DB, username string) (User, error) {
	var user User
	result := db.Where("username = ?", username).First(&user)
	return user, result.Error
}
