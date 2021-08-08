package server

import (
	"github.com/clauderoy790/gratitude-journal/helpers"
	http_helper "github.com/clauderoy790/gratitude-journal/http-helper"
	"net/http"
)

func (s *Server) loginHandler(writer http.ResponseWriter, request *http.Request) {
	params, err := http_helper.GetQueryParams(request)
	if err != nil {
		http_helper.WriteError(writer, err, http.StatusBadRequest)
		return
	}
	logRes := helpers.UserHelper.Login(params["email"], params["password"])
	http_helper.WriteJson(writer, logRes)
}
