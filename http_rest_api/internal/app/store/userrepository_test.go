package store_test

import (
	"github.com/gopherschool/http-rest-api/internal/app/models"
	"github.com/gopherschool/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
)

// tests for Create method
func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users") // cleaning users table

	u := models.TestUser(t)
	u.ID = strconv.Itoa(rand.Intn(1000))
	user, err := s.User().Create(u)
	assert.NoError(t, err) // check that no error raised
	assert.NotNil(t, user) // check that user is not nil
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users") // cleaning users table

	email := "user123@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := models.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByID(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u1, err := s.User().Create(models.TestUser(t))
	if err != nil {
		t.Fatal(err)
	}
	u2, err := s.User().FindByID(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
	assert.Equal(t, u2.ID, u1.ID)
}
