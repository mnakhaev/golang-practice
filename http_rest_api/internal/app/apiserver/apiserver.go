package apiserver

import (
	"github.com/gopherschool/http-rest-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// Starting new server with config
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("Starting API server...")

	// Using gorilla mux as router
	s.configureRouter()

	// Configuring DB store
	if err := s.configureStore(); err != nil {
		return err
	}

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	logrus.SetLevel(level)
	return nil
}

// configureRouter doesn't return any error because it describes the routing
func (s *APIServer) configureRouter() {
	// Creating new handleHello func to handle /hello endpoint
	s.router.HandleFunc("/hello", s.handleHello())
}

// configureStore ...
func (s *APIServer) configureStore() error {
	// Question: why not accessing from s.store from the beginning?
	// Why store.New ?
	st := store.New(s.config.Store)

	//Check for if err := s.store.Open(); err != nil {
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

// handleHello returns the whole interface instead of just func
func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello!")
	}
}
