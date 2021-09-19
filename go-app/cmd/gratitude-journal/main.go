package main

import (
	"context"
	"fmt"

	"github.com/clauderoy790/gratitude-journal/server"
)

func main() {
	s := server.New(context.Background())
	err := s.Run()
	if err != nil {
		panic(fmt.Sprintf("Server error occured: %v", err))
	}
}
