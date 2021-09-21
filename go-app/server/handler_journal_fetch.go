package server

import (
	"errors"
	"net/http"

	"github.com/clauderoy790/gratitude-journal/helper"
)

func (s *Server) journalFetchHandler(writer http.ResponseWriter, request *http.Request) {
	params, err := helper.ProcessJsonBody(request)
	if err != nil {
		helper.WriteError(writer, err, http.StatusBadRequest)
		return
	}
	userID := params["userID"]
	date := params["date"]

	if userID == "" || date == "" {
		helper.WriteError(writer, errors.New("must provide userID and date"), http.StatusBadRequest)
		return
	}
	//Read entry
	journalRes := helper.JournalHelper.GetEntry(userID, date)
	if journalRes.Error != "" {
		helper.WriteError(writer, errors.New(journalRes.Error), http.StatusInternalServerError)
	} else {
		helper.WriteJson(writer, journalRes.Entry)
	}
}
