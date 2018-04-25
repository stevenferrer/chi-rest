package memory

import (
	"errors"

	usermodel "github.com/moqafi/harper/model/user"
)

func New() usermodel.Storer {
	users := make([]usermodel.User, 0)
	return &Store{users: users}
}

// Store implements usermodel.Storer
type Store struct {
	users []usermodel.User
}

func (s *Store) List(filter ...usermodel.Filter) ([]usermodel.User, error) {
	return s.users, nil
}

func (s *Store) Get(id int64) (usermodel.User, error) {
	var user usermodel.User

	if !s.isUserIDInStore(id) {
		return user, errors.New("User id not found in store")
	}

	for _, u := range s.users {
		if u.ID == id {
			user = u
		}
	}

	return user, nil
}

// what's better Create, or Add?
func (s *Store) Create(u usermodel.User) error {
	if s.isUserEmailInStore(u.Email) {
		return errors.New("User email already in store")
	}

	s.users = append(s.users, u)

	return nil
}

func (s *Store) Update(u usermodel.User) error {
	return nil
}

func (s *Store) Delete(u usermodel.User) error {
	return nil
}

// Setup does the database setup for the User model
func (s *Store) Setup() error {
	return nil
}

// isUserIDInStore returns true if user id is already in the store
func (s *Store) isUserIDInStore(id int64) bool {
	for _, usr := range s.users {
		if usr.ID == id {
			return true
		}
	}

	return false
}

// isUserEmailInStore returns true if user email is already in the store
func (s *Store) isUserEmailInStore(email string) bool {
	for _, usr := range s.users {
		if usr.Email == email {
			return true
		}
	}

	return false
}
