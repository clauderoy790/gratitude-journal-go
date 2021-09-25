package repository

import "gorm.io/gorm"

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
	db    *gorm.DB
	user  User
	quote Quote
	entry JournalEntry
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetUser(email string) (*User, error) {
	r.db.Where("name = ?", "jinzhu").First(&user)
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
