package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	// Elo    string `json:"email" gorm:"unique"`
	// Password string `json:"password"` // hash in production
}
