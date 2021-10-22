package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/clauderoy790/gratitude-journal/helper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *Server) loginHandler(writer http.ResponseWriter, request *http.Request) {
	params, err := helper.ProcessJsonBody(request)
	if err != nil {
		helper.WriteError(writer, err, http.StatusBadRequest)
		return
	}

	email := params["email"]
	password := params["password"]

	user, err := s.repo.GetUser(email)
	if err == gorm.ErrRecordNotFound {
		helper.WriteError(writer, errors.New("this user is not registered"), http.StatusNotFound)
		return
	} else if err != nil {
		helper.WriteError(writer, fmt.Errorf("error: %w", err), http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		helper.WriteError(writer, errors.New("Password is not valid"), http.StatusUnauthorized)
		return
	}

	helper.WriteJson(writer, LoginResult{UserID: user.ID, Success: true})
}

type LoginResult struct {
	UserID  uint `json:"userId"`
	Success bool `json:"success"`
}
