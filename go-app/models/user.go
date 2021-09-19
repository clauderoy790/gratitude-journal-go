package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string    `json:"email,omitempty"`
	PasswordHash string    `json:"passwordHash,omitempty"`
	DateCreated  time.Time `json:"dateCreated,omitempty"`
}

type LoginResult struct {
	UserID  string `json:"userId"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type RegisterResult struct {
	UserID string `json:"userId"`
	Error  string `json:"error"`
}
