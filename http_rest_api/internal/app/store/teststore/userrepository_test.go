package teststore_test

import (
	"github.com/gopherschool/http-rest-api/internal/app/models"
	"github.com/gopherschool/http-rest-api/internal/app/store"
	"github.com/gopherschool/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
)

// tests for Create method
func TestUserRepository_Create(t *testing.T) {
	s := teststore.NewStore()
	u := models.TestUser(t)
	u.ID = strconv.Itoa(rand.Intn(1000))
	assert.NoError(t, s.User().Create(u)) // check that no error raised
	assert.NotNil(t, u)                   // check that user is not nil
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.NewStore()
	email := "user123@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := models.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
