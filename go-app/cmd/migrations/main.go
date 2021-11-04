package main

import (
	"fmt"

	"github.com/clauderoy790/gratitude-journal/config"
	"github.com/clauderoy790/gratitude-journal/repository"
)

func main() {
	cfg := config.Get()
	db, err := repository.ConnectToDatabase(&cfg)
	if err != nil {
		fmt.Println("failed to connect to the db: ", err)
		return
	}
	fmt.Println("successfully connected to the db!")
	err = db.AutoMigrate(&repository.User{}, &repository.Quote{}, &repository.JournalEntry{})
	if err != nil {
		fmt.Println("error migrating: ", err)
		return
	}
	fmt.Println("Migration successful!")
}
