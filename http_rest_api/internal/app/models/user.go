package models

// User models doesn't know anything about interaction with DB
// Repositories will be responsible for this kind of interaction
type User struct {
	ID                string
	Email             string
	EncryptedPassword string
}
