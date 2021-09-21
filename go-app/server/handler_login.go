package server

import (
	"net/http"

	"github.com/clauderoy790/gratitude-journal/helper"
)

func (s *Server) loginHandler(writer http.ResponseWriter, request *http.Request) {
	params, err := helper.ProcessJsonBody(request)
	if err != nil {
		helper.WriteError(writer, err, http.StatusBadRequest)
		return
	}
	logRes := helper.UserHelper.Login(params["email"], params["password"])
	helper.WriteJson(writer, logRes)
}
