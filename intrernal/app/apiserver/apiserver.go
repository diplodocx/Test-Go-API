package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func Start(s *APIServer) error {
	if err := ConfigureLogger(s); err != nil {
		return err
	}
	ConfigureRouter(s)
	s.logger.Info("started")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func ConfigureLogger(s *APIServer) error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	s.logger.SetLevel(level)
	return err
}

func ConfigureRouter(s *APIServer) {
	s.router.HandleFunc("/hello", HandleHello(s))
}

func HandleHello(s *APIServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}
