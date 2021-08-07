package server

import (
	"context"
	"github.com/clauderoy790/gratitude-journal/config"
	"github.com/clauderoy790/gratitude-journal/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	mux        *mux.Router
	ctx        context.Context
	cfg        *config.Config
	httpServer *http.Server
}

func New(ctx context.Context) *Server {
	server := Server{
		ctx: ctx,
	}
	server.setupRoutes()
	return &server
}

func (server *Server) Start() {
	helpers.MongoHelper.Connect()
	defer helpers.MongoHelper.Disconnect()

	server.httpServer = &http.Server{
		Handler: server.mux,
	}
}
