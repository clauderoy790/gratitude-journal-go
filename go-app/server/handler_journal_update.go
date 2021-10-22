package server

import (
	"encoding/json"
	"fmt"
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

	jEntry := repository.JournalEntry{}
	err = json.Unmarshal([]byte(params["entry"]), &jEntry)
	if err != nil {
		helper.WriteError(writer, fmt.Errorf("entry is not in a proper format: %w", err), http.StatusBadRequest)
		return
	}
	if err := s.repo.SaveJournalEntry(&jEntry); err != nil {
		helper.WriteError(writer, err, http.StatusInternalServerError)
		return
	}
	helper.WriteJson(writer, jEntry)
}
