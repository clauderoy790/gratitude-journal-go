package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) setupRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", s.homeHandler).Methods(http.MethodGet)
	r.HandleFunc("/health", s.healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/login", s.loginHandler).Methods(http.MethodPost)
	r.HandleFunc("/register", s.registerHandler).Methods(http.MethodPost)
	r.HandleFunc("/journal", s.journalUpdateHandler).Methods(http.MethodPut)
	r.HandleFunc("/journal", s.journalFetchHandler).Methods(http.MethodGet)
	r.Use(cors)
	s.muxRouter = r
}

func cors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		handler.ServeHTTP(w, r)
	})
}
