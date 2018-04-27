package memory

import (
	"errors"

	usermodel "github.com/moqafi/harper/model/user"
)

// New returns a new usermodel.Storer
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

func (s *Store) GetByID(id int64) (usermodel.User, error) {
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

func (s *Store) GetByEmail(email string) (usermodel.User, error) {
	var user usermodel.User

	if !s.isUserEmailInStore(email) {
		return user, errors.New("User email not found in store")
	}

	for _, u := range s.users {
		if u.Email == email {
			user = u
		}
	}

	return user, nil
}

// Create creates a new user
func (s *Store) Create(u usermodel.User) error {
	if s.isUserEmailInStore(u.Email) {
		return errors.New("User email already in store")
	}

	// assign a unique ID
	u.ID = int64(len(s.users)) + 1

	s.users = append(s.users, u)

	return nil
}

// Update updates an existing user
func (s *Store) UpdateByID(id int64, u usermodel.User) error {

	if !s.isUserIDInStore(id) {
		return errors.New("User not found in store")
	}

	idx := s.getUserIndexByID(id)

	s.users[idx] = u

	return nil
}

func (s *Store) UpdateByEmail(email string, u usermodel.User) error {
	if !s.isUserEmailInStore(email) {
		return errors.New("User not found in store")
	}

	idx := s.getUserIndexByEmail(email)

	s.users[idx] = u

	return nil
}

func (s *Store) Delete(u usermodel.User) error {
	if !s.isUserIDInStore(u.ID) {
		return errors.New("User not found in store")
	}

	idx := s.getUserIndexByID(u.ID)

	// delete item from slice trick
	s.users = append(s.users[:idx], s.users[idx+1:]...)

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

func (s *Store) getUserIndexByID(id int64) int64 {
	for idx, usr := range s.users {
		if usr.ID == id {
			return int64(idx)
		}
	}

	return -1
}

func (s *Store) getUserIndexByEmail(email string) int64 {
	for idx, usr := range s.users {
		if usr.Email == email {
			return int64(idx)
		}
	}

	return -1
}
