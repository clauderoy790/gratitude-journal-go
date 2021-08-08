package server

import (
	"github.com/clauderoy790/gratitude-journal/helpers"
	http_helper "github.com/clauderoy790/gratitude-journal/http-helper"
	"net/http"
)

func (s *Server) registerHandler(writer http.ResponseWriter, request *http.Request) {
	params, err := http_helper.ProcessJsonBody(request)
	if err != nil {
		http_helper.WriteError(writer, err, http.StatusBadRequest)
		return
	}

	regReg := helpers.UserHelper.Register(params["email"], params["password"], params["verifyPassword"])
	http_helper.WriteJson(writer, regReg)
}
