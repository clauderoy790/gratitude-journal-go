package server

import (
	"fmt"
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

	fmt.Sprintf("email:%v, password:%v, verify:%v\n", params["email"], params["password"], params["verifyPassword"])
	regReg := helpers.UserHelper.Register(params["email"].(string), params["password"].(string), params["verifyPassword"].(string))
	http_helper.WriteJson(writer, regReg)
}
