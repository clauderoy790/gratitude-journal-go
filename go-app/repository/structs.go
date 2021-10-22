package repository

import (
	"time"

	"gorm.io/gorm"
)

type JournalEntry struct {
	gorm.Model
	UserID         uint      `json:"userId,omitempty"`
	Date           time.Time `json:"date,omitempty"`
	Grateful1      string    `json:"grateful1,omitempty"`
	Grateful2      string    `json:"grateful2,omitempty"`
	Grateful3      string    `json:"grateful3,omitempty"`
	TodayGreat1    string    `json:"todayGreat1,omitempty"`
	TodayGreat2    string    `json:"todayGreat2,omitempty"`
	TodayGreat3    string    `json:"todayGreat3,omitempty"`
	Affirmation1   string    `json:"affirmation1,omitempty"`
	Affirmation2   string    `json:"affirmation2,omitempty"`
	HappenedGreat1 string    `json:"happenedGreat1,omitempty"`
	HappenedGreat2 string    `json:"happenedGreat2,omitempty"`
	HappenedGreat3 string    `json:"happenedGreat3,omitempty"`
	Better1        string    `json:"better1,omitempty"`
	Better2        string    `json:"better2,omitempty"`
	QuoteID        uint
	Quote          Quote `json:"quote,omitempty"`
}

type Quote struct {
	gorm.Model
	Message string `json:"message,omitempty"`
	Author  string `json:"author,omitempty"`
}

type User struct {
	gorm.Model
	Email        string `json:"email,omitempty"`
	PasswordHash string `json:"passwordHash,omitempty"`
}
