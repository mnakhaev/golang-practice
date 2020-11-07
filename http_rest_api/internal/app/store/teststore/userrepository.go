package teststore

import (
	"github.com/gopherschool/http-rest-api/internal/app/models"
	"github.com/gopherschool/http-rest-api/internal/app/store"
	"strconv"
)

// UserRepository structure for tests
type UserRepository struct {
	store *Store
	users map[string]*models.User
}

// Create test user in `users` map
func (r *UserRepository) Create(u *models.User) error {
	// check if user is valid. if OK - run BeforeCreate callback
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = strconv.Itoa(len(r.users))

	return nil
}

// FindByEmail in `users` map
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

// TODO: implement till the end
// FindByID in `users` map
func (r *UserRepository) FindByID(ID string) (*models.User, error) {
	u, ok := r.users[ID]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}
