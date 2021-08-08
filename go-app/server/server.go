package server

import (
	"context"
	"fmt"
	"github.com/clauderoy790/gratitude-journal/config"
	"github.com/clauderoy790/gratitude-journal/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	muxRouter  *mux.Router
	ctx        context.Context
	cfg        config.Config
	httpServer *http.Server
}

type M map[string]interface{}

func New(ctx context.Context) *Server {
	server := Server{
		ctx: ctx,
		cfg: config.Get(),
	}
	server.setupRoutes()
	server.httpServer = &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", server.cfg.App.Port),
		Handler: server.muxRouter,
	}
	return &server
}

func (s *Server) Start() error {
	helpers.MongoHelper.Connect()
	defer helpers.MongoHelper.Disconnect()

	fmt.Printf("Server started server on port: %d\n", s.cfg.App.Port)
	err := s.HttpServer().ListenAndServe()
	return err
}

func (s *Server) HttpServer() *http.Server {
	return s.httpServer
}
