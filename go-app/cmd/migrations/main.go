package main

import (
	"fmt"

	"github.com/clauderoy790/gratitude-journal/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:admin@tcp(127.0.0.1:3307)/daily_gratitude?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
