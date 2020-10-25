package store

import (
	"database/sql"

	_ "github.com/lib/pq" // Anonymous import
)

type Store struct {
	config         *Config
	db             *sql.DB
	UserRepository *UserRepository
}

// New returns pointer on store (not Config as in example)
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open will be used for store initialization
func (s *Store) Open() error {
	// sql.Open initializes 'lazy' connection to DB
	// Real connection is established during first call to DB
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	// db.Ping() makes simple test SELECT query to database
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db

	return nil
}

// Close will be used to disconnect from DB and perform other operations
func (s *Store) Close() {
	s.db.Close()
}

// User is special method to avoid using repositories without the store.
// Example of such call: store.User().Create()
func (s *Store) User() *UserRepository {
	if s.UserRepository != nil {
		return nil
	}

	s.UserRepository = &UserRepository{store: s}
	return s.UserRepository
}
