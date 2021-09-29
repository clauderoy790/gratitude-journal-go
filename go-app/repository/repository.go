package repository

import (
	"fmt"
	"time"

	"github.com/clauderoy790/gratitude-journal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository interface {
	GetUser(email string) (*User, error)
	CreateUser(email, passwordHash string) error
	DeleteUser(email string) error

	SaveJournalEntry(entry *JournalEntry) error
	GetJournalEntry(userID uint, date time.Time) (*JournalEntry, error)
	DeleteJournalEntry(userID uint, date time.Time) error

	QuotesCount() (int, error)
	GetQuote(id uint) (Quote, error)
	SaveQuote(message, author string) (uint, error)
	DeleteQuote(id uint) error
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

func (r *repository) CreateUser(email, passwordHash string) error {
	user := User{Email: email, PasswordHash: passwordHash}
	if r.db.Where("email = ?", email).First(&User{}).Error == gorm.ErrRecordNotFound {
		return r.db.Create(&user).Error
	}
	return fmt.Errorf("user with email %v already exists", email)
}

func (r *repository) DeleteUser(email string) error {
	return r.db.Unscoped().Where("email = ?", email).Delete(&User{}).Error
}

func (r *repository) SaveJournalEntry(entry *JournalEntry) error {
	if entry == nil {
		return fmt.Errorf("Cannot save nil entry")
	}
	dbEntry := JournalEntry{}

	where := r.db.Model(&JournalEntry{}).Where("user_id = ? AND date = ?", entry.UserID, entry.Date)
	result := where.First(&dbEntry)
	if result.Error == gorm.ErrRecordNotFound {
		return r.db.Create(entry).Error
	} else if result.RowsAffected == 1 {
		return where.Updates(entry).Error
	}
	return result.Error
}

func (r *repository) GetJournalEntry(userID uint, date time.Time) (*JournalEntry, error) {
	entry := JournalEntry{
		UserID: userID,
		Date:   r.roundDate(date),
	}

	err := r.db.Preload("Quote").Where("user_id = ? AND date = ?", userID, entry.Date).First(&entry).Error
	if err == gorm.ErrRecordNotFound {
		err = r.db.Create(&entry).Error
	}
	return &entry, err
}

func (r *repository) DeleteJournalEntry(userID uint, date time.Time) error {
	formattedDate := r.roundDate(date)
	return r.db.Unscoped().Where("user_id = ? AND date = ?", userID, formattedDate).Delete(&JournalEntry{}).Error
}

func (r *repository) roundDate(date time.Time) time.Time {
	formatted := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	return formatted
}

func (r *repository) GetQuotes() ([]Quote, error) {
	var quotes []Quote
	err := r.db.Find(&quotes).Error
	return quotes, err
}

func (r *repository) SaveQuote(message, author string) (uint, error) {
	quote := Quote{
		Message: message,
		Author:  author,
	}
	err := r.db.Create(&quote).Error
	if err != nil {
		return 0, fmt.Errorf("error creating quote: %w", err)
	}
	return quote.ID, nil
}

func (r *repository) DeleteQuote(id uint) error {
	return r.db.Unscoped().Delete(&Quote{}, id).Error
}

func (r *repository) DeleteAllQuotes() error {
	return r.db.Unscoped().Where("1 = 1").Delete(&Quote{}).Error
}

func (r *repository) QuotesCount() (int, error) {
	var count int64
	err := r.db.Model(&Quote{}).Count(&count).Error
	return int(count), err
}

func (r *repository) GetQuote(id uint) (Quote, error) {
	var quote Quote
	err := r.db.Where("id = ?", id).First(&quote).Error
	return quote, err
}

func ConnectToDatabase(cfg *config.Config) (*gorm.DB, error) {
	db := cfg.Database
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", db.User, db.Password, db.Host, db.Port, db.Name)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
