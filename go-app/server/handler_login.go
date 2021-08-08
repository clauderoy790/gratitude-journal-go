package server

import (
	"github.com/clauderoy790/gratitude-journal/helpers"
	http_helper "github.com/clauderoy790/gratitude-journal/http-helper"
	"net/http"
)

func (s *Server) loginHandler(writer http.ResponseWriter, request *http.Request) {
	params, err := http_helper.ProcessJsonBody(request)
	if err != nil {
		http_helper.WriteError(writer, err, http.StatusBadRequest)
		return
	}
	logRes := helpers.UserHelper.Login(params["email"].(string), params["password"].(string))
	http_helper.WriteJson(writer, logRes)
}
