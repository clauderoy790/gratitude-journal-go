package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/clauderoy790/gratitude-journal/helper"
	"github.com/clauderoy790/gratitude-journal/repository"
)

func (s *Server) journalUpdateHandler(writer http.ResponseWriter, request *http.Request) {
	params, err := helper.ProcessJsonBody(request)
	if err != nil {
		helper.WriteError(writer, err, http.StatusBadRequest)
		return
	}
	userID := params["userID"]
	date := params["date"]
	entry := params["entry"]

	if userID == "" || date == "" {
		helper.WriteError(writer, errors.New("must provide userID, date and entry"), http.StatusBadRequest)
		return
	}
	jEntry := repository.JournalEntry{}
	err = json.Unmarshal([]byte(entry), &jEntry)
	if err != nil {
		helper.WriteError(writer, errors.New("entry is not in a valid format"), http.StatusBadRequest)
	} else {
		if err := helper.JournalHelper.WriteEntry(userID, date, jEntry); err != nil {
			helper.WriteError(writer, err, http.StatusInternalServerError)
		} else {
			helper.WriteJson(writer, jEntry)
		}
	}
}
