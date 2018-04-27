package mssql

import (
	"database/sql"

	usermodel "github.com/moqafi/harper/model/user"
)

func New(db *sql.DB) usermodel.Storer {
	return nil
}

type Store struct {
	db *sql.DB
}

func (s *Store) List(filter ...usermodel.Filter) ([]usermodel.User, error) {
	return nil, nil
}

func (s *Store) GetByID(id int64) (usermodel.User, error) {
	return usermodel.User{}, nil
}

func (s *Store) GetByEmail(email string) (usermodel.User, error) {
	return usermodel.User{}, nil
}

func (s *Store) Create(user usermodel.User) error {
	return nil
}

func (s *Store) UpdateByID(id int64, user usermodel.User) error {
	return nil
}

func (s *Store) UpdateByEmail(email string, user usermodel.User) error {
	return nil
}

func (s *Store) Delete(user usermodel.User) error {
	return nil
}
