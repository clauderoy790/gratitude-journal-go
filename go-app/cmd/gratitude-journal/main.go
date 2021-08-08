package main

import (
	"context"
	"fmt"
	"github.com/clauderoy790/gratitude-journal/server"
)

func main() {
	s := server.New(context.Background())
	err := s.Start()
	if err != nil {
		fmt.Println("A server error occurred: ", err)
	}
}
