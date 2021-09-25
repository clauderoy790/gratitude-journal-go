package main

import (
	"context"
	"fmt"

	"github.com/clauderoy790/gratitude-journal/config"
	"github.com/clauderoy790/gratitude-journal/repository"
	"github.com/clauderoy790/gratitude-journal/server"
)

func main() {
	cfg := config.Get()
	db, err := connectToDatabase(cfg.Database)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %w", err))
	}

	repo := repository.NewRepository(db)
	s := server.New(context.Background(), repo)
	err := s.Run()
	if err != nil {
		panic(fmt.Sprintf("Server error occured: %w", err))
	}
}

func connectToDatabase() {

}
