package main

import (
	"context"
	"github.com/clauderoy790/gratitude-journal/server"
)

func main() {
	s := server.New(context.Background())
	s.Start()
}
