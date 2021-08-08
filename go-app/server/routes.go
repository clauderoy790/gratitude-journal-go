package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s *Server) setupRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", s.homeHandler).Methods(http.MethodGet)
	r.HandleFunc("/login", s.loginHandler).Methods(http.MethodPost)
	r.HandleFunc("/register", s.registerHandler).Methods(http.MethodPost)
	r.HandleFunc("/journal", s.journalUpdateHandler).Methods(http.MethodPut)
	r.HandleFunc("/journal", s.journalFetchHandler).Methods(http.MethodGet)
	s.muxRouter = r
}
