package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/gopherschool/http-rest-api/internal/app/models"
	"github.com/gopherschool/http-rest-api/internal/app/store/teststore"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerHandleUsersCreate(t *testing.T) {
	s := newServer(teststore.NewStore(), sessions.NewCookieStore([]byte("random_secret")))
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid data",
			payload: map[string]string{
				"email":    "user@example.org",
				"password": "password",
			},
			expectedCode: http.StatusCreated, // see handleUsersCreate() in server.go
		},
		{
			name:         "invalid payload",
			payload:      "invalid_payload",
			expectedCode: http.StatusBadRequest, // see handleUsersCreate() in server.go
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "invalid",
			},
			expectedCode: http.StatusUnprocessableEntity, // see handleUsersCreate() in server.go
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// TODO: find out about `httptest.NewRecorder`
			rec := httptest.NewRecorder()

			// TODO: find out about next 2 lines
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)

			// Perform new request to /users endpoint with payload defined above
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			// TODO: find out about `ServeHTTP`
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServerHandleSessionsCreate(t *testing.T) {
	u := models.TestUser(t)
	store := teststore.NewStore()
	store.User().Create(u)
	s := newServer(store, sessions.NewCookieStore([]byte("random_secret")))
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid data",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "invalid",
				"password": "qwe123QWE",
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    "invalid",
				"password": "qwe123QWE!@#",
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)

			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
