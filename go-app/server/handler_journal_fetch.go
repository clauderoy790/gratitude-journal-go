package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/clauderoy790/gratitude-journal/helper"
	"github.com/clauderoy790/gratitude-journal/repository"
)

func (s *Server) journalFetchHandler(writer http.ResponseWriter, request *http.Request) {
	params, err := helper.ProcessJsonBody(request)
	if err != nil {
		helper.WriteError(writer, err, http.StatusBadRequest)
		return
	}
	userID, err := strconv.ParseUint(params["userID"], 10, 32)
	if err != nil {
		helper.WriteError(writer, fmt.Errorf("error parsing userID: %w", err), http.StatusBadRequest)
		return
	}

	date, err := time.Parse("", params["date"])
	if err != nil {
		helper.WriteError(writer, fmt.Errorf("error parsing date", err), http.StatusBadRequest)
		return
	}

	//Read entry
	entry, err := s.repo.GetJournalEntry(uint(userID), date)
	if err != nil {
		helper.WriteError(writer, fmt.Errorf("error getting entry: %w", err), http.StatusInternalServerError)
		return
	}
	helper.WriteJson(writer, entry)
}

type JournalEntryResponse struct {
	Entry repository.JournalEntry `json:"entry"`
	Error string                  `json:"error"`
}

type JournalEntryRequest struct {
	Date   string                  `json:"date"`
	UserID string                  `json:"userID"`
	Entry  repository.JournalEntry `json:"entry"`
}
