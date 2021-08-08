package server

import (
	http_helper "github.com/clauderoy790/gratitude-journal/http-helper"
	"net/http"
)

func (s *Server) homeHandler(writer http.ResponseWriter, request *http.Request) {
	http_helper.WriteJson(writer, M{
		"message": "welcome to daily gratitude",
	})
}
