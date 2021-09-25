package repository

import (
	"fmt"

	"github.com/clauderoy790/gratitude-journal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository interface {
	GetUser(email string) (*User, error)
	SaveUser(user *User) error

	SaveJournalEntry(entry *JournalEntry) error
	DeleteJournalEntry(entry *JournalEntry) error
	GetJournalEntry(userID, date string) (*JournalEntry, error)

	GetQuotes() ([]*Quote, error)
	SaveQuote(quote *Quote) error
	DeleteAllQuotes() error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetUser(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *repository) SaveUser(user *User) error {
	panic("not implemented") // TODO: Implement
}

func (r *repository) SaveJournalEntry(entry *JournalEntry) error {
	panic("not implemented") // TODO: Implement
}

func (r *repository) DeleteJournalEntry(entry *JournalEntry) error {
	panic("not implemented") // TODO: Implement
}

func (r *repository) GetJournalEntry(userID string, date string) (*JournalEntry, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repository) GetQuotes() ([]*Quote, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repository) SaveQuote(quote *Quote) error {
	panic("not implemented") // TODO: Implement
}

func (r *repository) DeleteAllQuotes() error {
	panic("not implemented") // TODO: Implement
}

func ConnectToDatabase(cfg *config.Config) (*gorm.DB, error) {
	db := cfg.Database
	fmt.Println("db config: ", db)
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", db.User, db.Password, db.Host, db.Port, db.Name)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
