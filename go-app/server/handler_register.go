package server

import (
	"encoding/json"
	"fmt"
	"github.com/clauderoy790/gratitude-journal/helpers"
	http_helper "github.com/clauderoy790/gratitude-journal/http-helper"
	"net/http"
)

func (s *Server) registerHandler(writer http.ResponseWriter, request *http.Request) {
	params := make(map[string]string)
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&params)
	if err != nil {
		fmt.Println("error decoding json: ", err)
		http_helper.WriteError(writer, err, http.StatusBadRequest)
		return
	}

	fmt.Sprintf("email:%v, password:%v, verify:%v\n", params["email"], params["password"], params["verifyPassword"])
	regReg := helpers.UserHelper.Register(params["email"], params["password"], params["verifyPassword"])
	http_helper.WriteJson(writer, regReg)
}
