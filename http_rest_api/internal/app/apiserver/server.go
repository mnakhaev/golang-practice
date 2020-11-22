package apiserver

import (
	"encoding/json"
	"errors"
	"github.com/gopherschool/http-rest-api/internal/app/models"
	"github.com/gopherschool/http-rest-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

// handle incoming requests and realize HTTP handler interface

const (
	sessionName = "simple_session_name" // will be returned as response cookie - Set-Cookie: simple_session_name=MTN..
)

// create common error for both wrong email and password (more secire way)
var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
)

type server struct {
	router       *mux.Router
	store        store.Store    // it's an interface
	sessionStore sessions.Store // gorilla session. Will be returned as response cookie
}

// newServer accepts store interface
func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		store:        store,
		sessionStore: sessionStore,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	// Create new session for user. Will be returned as response header
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")
}

// Check if this function returns pointer
func (s *server) handleUsersCreate() http.HandlerFunc {
	// request describes parameters needed for authentication of user
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			// error accepts response writer, request and response code (400) and error
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &models.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := s.store.User().Create(u); err != nil {
			// User send incorrect data - 422 error
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		// hide password and render user without it
		u.Sanitize()
		// since user `u` is passed to `respond` method, need to set JSON tags in base User struct
		s.respond(w, r, http.StatusCreated, u)

	}
}

func (s *server) handleSessionsCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		// find user by email and check that email is OK and passed password corresponds to encrypted one in store
		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePasswords(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}
		// return cookie to user after successful authentication
		// using gorilla/sessions package for that
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			// return internal server error because problem is on our side
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		// Load session if it has user ID. Otherwise, return unauthorized error or smth else
		session.Values["user_id"] = u.ID
		// Saving current session
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

// error is helper method to render any errors during work of handlers
// it will use another helper named `respond`
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

// respond is used for rendering of everything. `data` can have any type - set empty interface
func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
