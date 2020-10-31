package store

import (
	"fmt"
	"strings"
	"testing"
)

// helper module needed for tests
// return test store with required configuration
// and return function which will cleanup the tables
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper() // TODO: read about this method

	config := NewConfig()
	config.DatabaseURL = databaseURL
	store := NewStore(config) // creating new store

	if err := store.Open(); err != nil {
		t.Fatal(err)
	}

	// returning store and teardown function which will delete some tables
	return store, func(tables ...string) {
		// number of tables can be zero
		if len(tables) > 0 {
			// TODO: check that
			if _, err := store.db.Exec(
				fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")),
			); err != nil {
				t.Fatal(err)
			}
		}

		store.Close()
	}
}
