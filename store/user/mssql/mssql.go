package mssql

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	usermodel "github.com/moqafi/harper/model/user"
)

func New(db *gorm.DB) usermodel.Storer {
	return &Store{db}
}

type Store struct {
	db *gorm.DB
}

func (s *Store) List(filter ...usermodel.Filter) ([]usermodel.User, error) {
	// return all users for now
	var users []usermodel.User
	err := s.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Store) GetByID(id uint64) (usermodel.User, error) {

	return usermodel.User{}, nil
}

func (s *Store) GetByEmail(email string) (usermodel.User, error) {
	var user usermodel.User

	err := s.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (s *Store) Create(user usermodel.User) (usermodel.User, error) {

	hash, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = hash
	//	var user usermodel.User
	if err := s.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (s *Store) UpdateByID(id uint64, user usermodel.User) (usermodel.User, error) {
	//	var user usermodel.User
	return user, nil
}

func (s *Store) UpdateByEmail(email string, user usermodel.User) (usermodel.User, error) {
	//	var user usermodel.User
	return user, nil
}

func (s *Store) Delete(user usermodel.User) (usermodel.User, error) {
	//	var user usermodel.User
	return user, nil
}
