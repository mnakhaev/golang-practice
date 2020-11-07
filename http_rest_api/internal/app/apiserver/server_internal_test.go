package apiserver

import (
	"github.com/bmizerany/assert"
	"github.com/gopherschool/http-rest-api/internal/app/store/teststore"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerHandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s := newServer(teststore.NewStore())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
