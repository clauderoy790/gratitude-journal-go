package server

import (
	"github.com/clauderoy790/gratitude-journal/helper"
	"net/http"
)

func (s *Server) registerHandler(writer http.ResponseWriter, request *http.Request) {
	params, err := helper.ProcessJsonBody(request)
	if err != nil {
		helper.WriteError(writer, err, http.StatusBadRequest)
		return
	}

	regReg := helper.UserHelper.Register(params["email"], params["password"], params["verifyPassword"])
	helper.WriteJson(writer, regReg)
}
