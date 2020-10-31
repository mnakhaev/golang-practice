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

	u, err := s.User().Create(&models.User{
		ID:                strconv.Itoa(rand.Intn(1000)),
		Email:             "user@example.org",
		EncryptedPassword: "qwe123QWE",
	})
	assert.NoError(t, err) // check that no error raised
	assert.NotNil(t, u)    // check that user is not nil
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users") // cleaning users table

	email := "user123@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&models.User{Email: email})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByID(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u1, err := s.User().Create(&models.User{Email: "test_user@example.org"})
	if err != nil {
		t.Fatal(err)
	}
	u2, err := s.User().FindByID(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
	assert.Equal(t, u2.ID, u1.ID)
}
