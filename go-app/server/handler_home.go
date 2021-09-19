package server

import (
	"github.com/clauderoy790/gratitude-journal/helper"
	"net/http"
)

func (s *Server) homeHandler(writer http.ResponseWriter, request *http.Request) {
	helper.WriteJson(writer, M{
		"message": "welcome to daily gratitude",
	})
}
