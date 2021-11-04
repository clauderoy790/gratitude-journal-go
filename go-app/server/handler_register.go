package server

import (
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/clauderoy790/gratitude-journal/helper"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) registerHandler(writer http.ResponseWriter, request *http.Request) {
	params, err := helper.ProcessJsonBody(request)
	if err != nil {
		helper.WriteError(writer, err, http.StatusBadRequest)
		return
	}

	email := params["email"]
	password := params["password"]
	verifyPassword := params["verifyPassword"]

	if !isValidEmail(email) {
		helper.WriteError(writer, errors.New("Email must be valid."), http.StatusBadRequest)
		return
	}
	if len(password) < s.cfg.App.MinPasswordLength {
		helper.WriteError(writer, fmt.Errorf("Password must be at least %v characters.", s.cfg.App.MinPasswordLength), http.StatusBadRequest)
		return
	}
	if password != verifyPassword {
		helper.WriteError(writer, errors.New("Passwords are not identical."), http.StatusBadRequest)
		return
	}

	hashedPass, err := hashPass(password)
	if err != nil {
		helper.WriteError(writer, errors.New("Error creating user"), http.StatusInternalServerError)
		return
	}

	userID, err := s.repo.CreateUser(email, hashedPass)
	if err != nil {
		helper.WriteError(writer, fmt.Errorf("Failed to get user with email %s: %w", email, err), http.StatusNotFound)
	}

	helper.WriteJson(writer, RegisterResult{
		UserID: userID,
	})
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func hashPass(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

type RegisterResult struct {
	UserID uint `json:"userId"`
}
