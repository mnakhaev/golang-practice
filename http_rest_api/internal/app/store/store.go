package store

type Store struct {
	config *Config
}

func New(config *Config) *Store {
	return &Config{
		config: config,
	}
}

// Open will be used for store initialization
func (s *Store) Open() error {
	return nil
}

// Close will be used to disconnect from DB and perform other operations
func (s *Store) Close() {

}
