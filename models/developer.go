package models

import (
	"github.com/prasad89/devspace-api/initializers"
)

type Developer struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

func GetByUsername(username string) (Developer, error) {
	var developer Developer
	result := initializers.DB.Where("username = ?", username).First(&developer)
	return developer, result.Error
}
