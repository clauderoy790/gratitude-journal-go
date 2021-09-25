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
	db, err := repository.ConnectToDatabase(&cfg)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}

	repo := repository.NewRepository(db)
	s := server.New(context.Background(), repo, cfg)
	err = s.Run()
	if err != nil {
		panic(fmt.Sprintf("Server error occured: %v", err))
	}
}
