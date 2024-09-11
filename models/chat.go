package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Message  string `json:"message"`
}
