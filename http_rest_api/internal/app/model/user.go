package model

// User model doesn't know anything about interaction with DB
// Repositories will be responsible for this kind of interaction
type User struct {
	ID                string
	Email             string
	EnctyptedPassword string
}
