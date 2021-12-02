package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/clauderoy790/gratitude-journal/config"
	"github.com/clauderoy790/gratitude-journal/repository"
	"github.com/clauderoy790/gratitude-journal/server"
)

func main() {
	verPtr := flag.String("version", "", "version of the app")
	commitHashPtr := flag.String("commit-hash", "", "current commit hash")
	flag.Parse()

	if *verPtr == "" || *commitHashPtr == "" {
		panic(fmt.Sprintf("Invalid version or commit hash. Version: %v, commit hash: %v", *verPtr, *commitHashPtr))
	}

	config.VERSION = *verPtr
	config.COMMITHASH = *commitHashPtr

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
