package server

import (
	"net/http"

	"github.com/clauderoy790/gratitude-journal/config"
	"github.com/clauderoy790/gratitude-journal/helper"
)

func (s *Server) healthHandler(writer http.ResponseWriter, request *http.Request) {
	helper.WriteJson(writer, health{
		Version:    config.VERSION,
		CommitHash: config.COMMITHASH,
	})
}

type health struct {
	Version    string `json:"version"`
	CommitHash string `json:"commitHash"`
}
