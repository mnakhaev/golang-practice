package store

import "github.com/gopherschool/http-rest-api/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	// postgres doesn't return IDs by default, but we need to get this ID for successfully created user
	// this ID will be used later somehow
	// Scan method is used to map returned string to passed arguments (should be pointers!)
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EnctyptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

// FindByEmail method is needed for authorization
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
