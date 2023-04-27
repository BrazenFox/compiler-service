package server

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRuoter()

	s.logger.Info("starting server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Server) configureRuoter(){
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *Server) handleHello() http.HandlerFunc{
	type request struct{
		name string
	}

	return func (w http.ResponseWriter, r *http.Request)  {
		io.WriteString(w, "Hello")
	}
}
