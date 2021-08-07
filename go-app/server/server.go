package server

import (
	"context"
	"fmt"
	"github.com/clauderoy790/gratitude-journal/config"
	"github.com/clauderoy790/gratitude-journal/helpers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	muxRouter  *mux.Router
	ctx        context.Context
	cfg        config.Config
	httpServer *http.Server
}

func New(ctx context.Context) *Server {
	server := Server{
		ctx: ctx,
		cfg: config.Get(),
	}
	server.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", server.cfg.App.Port),
		Handler: server.muxRouter,
	}
	server.setupRoutes()
	return &server
}

func (s *Server) Start() {
	helpers.MongoHelper.Connect()
	defer helpers.MongoHelper.Disconnect()

	fmt.Printf("Server started server on port: %d\n", s.cfg.App.Port)
	err := s.HttpServer().ListenAndServe()
	if err != nil {
		log.Fatalln("s unexpectedly stopped", err)
	}
}

func (s *Server) HttpServer() *http.Server {
	return s.httpServer
}
