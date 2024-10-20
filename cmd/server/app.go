package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
)

type Server struct {
	logger *slog.Logger
	Router *gin.Engine
}

func (s *Server) init() *Server {
	s.Router = gin.Default()
	s.logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return s
}

func (s *Server) Run(addr string) {
	if err := http.ListenAndServe(addr, s.Router); err != nil {
		s.logger.Error(err.Error())
	}
}

func (s *Server) InitializeRedis() *Server {
	store, err := redis.NewStore(10, "tcp",
		os.Getenv("PUBLIC_HOST")+os.Getenv("REDIS_PORT"),
		"", []byte("secret"))
	if err != nil {
		s.logger.Error(err.Error())
	} else {
		s.logger.Info("Redis session initialized")
	}
	s.Router.Use(sessions.Sessions("sessions", store))

	return s
}
