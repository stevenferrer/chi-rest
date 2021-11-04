package memory

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	usermodel "github.com/stevenferrer/chi-rest/model/user"
)

// New returns a new usermodel.Storer
func New() *Store {
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

func (s *Store) GetByID(id uint64) (usermodel.User, error) {
	var user usermodel.User

	if !s.isUserIDInStore(id) {
		return user, errors.New("User not found in store")
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
		return user, errors.New("User not found in store")
	}

	for _, u := range s.users {
		if u.Email == email {
			user = u
		}
	}

	return user, nil
}

// Create creates a new user
func (s *Store) Create(u usermodel.User) (usermodel.User, error) {
	var user usermodel.User
	if err := u.Validate(); err != nil {
		return user, err
	}

	if s.isUserEmailInStore(u.Email) {
		return user, errors.New("User email already in store")
	}

	// assign a unique ID
	u.ID = uint64(len(s.users)) + 1
	hash, err := bcrypt.GenerateFromPassword(u.Password, bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	u.Password = hash

	s.users = append(s.users, u)

	user = u

	return user, nil
}

// Update updates an existing user
func (s *Store) UpdateByID(id uint64, u usermodel.User) (usermodel.User, error) {
	var user usermodel.User

	if !s.isUserIDInStore(id) {
		return user, errors.New("User not found in store")
	}

	idx := s.getUserIndexByID(id)

	hash, err := bcrypt.GenerateFromPassword(u.Password, bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	u.Password = hash

	s.users[idx] = u

	user = u

	return user, nil
}

func (s *Store) UpdateByEmail(email string, u usermodel.User) (usermodel.User, error) {
	var user usermodel.User

	if !s.isUserEmailInStore(email) {
		return user, errors.New("User not found in store")
	}

	idx := s.getUserIndexByEmail(email)

	hash, err := bcrypt.GenerateFromPassword(u.Password, bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	u.Password = hash

	s.users[idx] = u

	user = u

	return user, nil
}

func (s *Store) Delete(u usermodel.User) (usermodel.User, error) {
	var user usermodel.User
	if !s.isUserIDInStore(u.ID) {
		return user, errors.New("User not found in store")
	}

	idx := s.getUserIndexByID(u.ID)

	// get user to be deleted
	user = s.users[idx]

	// delete item from slice trick
	s.users = append(s.users[:idx], s.users[idx+1:]...)

	return user, nil
}

// isUserIDInStore returns true if user id is already in the store
func (s *Store) isUserIDInStore(id uint64) bool {
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

func (s *Store) getUserIndexByID(id uint64) uint64 {
	for idx, usr := range s.users {
		if usr.ID == id {
			return uint64(idx)
		}
	}

	return 0
}

func (s *Store) getUserIndexByEmail(email string) uint64 {
	for idx, usr := range s.users {
		if usr.Email == email {
			return uint64(idx)
		}
	}

	return 0
}
