package models

import (
	"gorm.io/gorm"
)

type Quote struct {
	gorm.Model
	Message string `json:"message,omitempty"`
	Author  string `json:"author,omitempty"`
}

type QuoteResult struct {
	Message string `json:"message"`
	Author  string `json:"author"`
	Error   string `json:"error"`
}
