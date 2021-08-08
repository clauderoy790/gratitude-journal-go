package server

import (
	"errors"
	"github.com/clauderoy790/gratitude-journal/helpers"
	http_helper "github.com/clauderoy790/gratitude-journal/http-helper"
	"net/http"
)

func (s *Server) journalFetchHandler(writer http.ResponseWriter, request *http.Request) {
	params, err := http_helper.ProcessJsonBody(request)
	if err != nil {
		http_helper.WriteError(writer, err, http.StatusBadRequest)
		return
	}
	userID := params["userID"].(string)
	date := params["date"].(string)

	if userID == "" || date == "" {
		http_helper.WriteError(writer, errors.New("must provide userID and date"), http.StatusBadRequest)
		return
	}
	//Read entry
	journalRes := helpers.JournalHelper.GetEntry(userID, date)
	if journalRes.Error != "" {
		http_helper.WriteError(writer, errors.New(journalRes.Error), http.StatusInternalServerError)
	} else {
		http_helper.WriteJson(writer, journalRes.Entry)
	}
}
