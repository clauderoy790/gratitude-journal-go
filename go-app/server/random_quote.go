package server

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"

	"github.com/clauderoy790/gratitude-journal/repository"
)

func (s *Server) GetRandomQuote(userID, date string) (*repository.Quote, error) {
	mixedId, err := combineToStr(userID, date, 7)
	if err != nil {
		return nil, err
	}
	s1 := rand.NewSource(int64(mixedId))
	r1 := rand.New(s1)

	count, err := s.repo.QuotesCount()

	if err != nil {
		return nil, fmt.Errorf("Failed to get quotes count: %w", err)
	}

	firstID, err := s.repo.GetFirstQuoteID()
	if err != nil {
		return nil, fmt.Errorf("error getting first quote: %w", err)
	}

	randNb := r1.Intn(count) + int(firstID)
	randID := uint(randNb)
	quote, err := s.repo.GetQuote(randID)
	if err != nil {
		return nil, err
	}
	return &quote, nil
}

func combineToStr(id, date string, maxLen int) (nb int, err error) {

	// Make a Regex to say we only want numbers
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return 0, err
	}

	// Remove all characters for combined Id + Date
	idNumbers := reg.ReplaceAllString(id, "")
	idNumbers = clamp(idNumbers, maxLen)
	processedString := idNumbers + reg.ReplaceAllString(date, "")

	nb, err = strconv.Atoi(processedString)
	return nb, err
}

func clamp(str string, max int) string {
	if len(str) > max {
		str = str[:max]
	} else {
		str = str[:len(str)]
	}
	return str
}
