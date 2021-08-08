package server

import (
	"encoding/json"
	"errors"
	"github.com/clauderoy790/gratitude-journal/helpers"
	http_helper "github.com/clauderoy790/gratitude-journal/http-helper"
	"github.com/clauderoy790/gratitude-journal/models"
	"net/http"
)

func (s *Server) journalUpdateHandler(writer http.ResponseWriter, request *http.Request) {
	params, err := http_helper.ProcessJsonBody(request)
	if err != nil {
		http_helper.WriteError(writer, err, http.StatusBadRequest)
		return
	}
	userID := params["userID"]
	date := params["date"]
	entry := params["entry"]

	if userID == "" || date == "" {
		http_helper.WriteError(writer, errors.New("must provide userID, date and entry"), http.StatusBadRequest)
		return
	}
	jEntry := models.JournalEntry{}
	err = json.Unmarshal([]byte(entry), &jEntry)
	if err != nil {
		http_helper.WriteError(writer, errors.New("entry is not in a valid format"), http.StatusBadRequest)
	} else {
		if err := helpers.JournalHelper.WriteEntry(userID, date, jEntry); err != nil {
			http_helper.WriteError(writer, err, http.StatusInternalServerError)
		} else {
			http_helper.WriteJson(writer, jEntry)
		}
	}
}
