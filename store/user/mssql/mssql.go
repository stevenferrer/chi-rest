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
	var user usermodel.User
	err := s.db.Where("id = ?", id).First(&user).Error
	return user, err
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

func (s *Store) UpdateByID(id uint64, newUser usermodel.User) (usermodel.User, error) {
	var user usermodel.User
	user, err := s.GetByID(id)
	if err != nil {
		return user, err
	}

	// updated password hash
	hash, err := bcrypt.GenerateFromPassword(
		newUser.Password,
		bcrypt.DefaultCost,
	)
	if err != nil {
		return user, err
	}
	newUser.Password = hash

	err = s.db.Model(&user).Updates(
		usermodel.User{
			Email:    newUser.Email,
			Password: hash,
		},
	).Error

	return user, err
}

func (s *Store) UpdateByEmail(email string, newUser usermodel.User) (usermodel.User, error) {
	var user usermodel.User
	user, err := s.GetByEmail(email)
	if err != nil {
		return user, err
	}

	// updated password hash
	hash, err := bcrypt.GenerateFromPassword(
		newUser.Password,
		bcrypt.DefaultCost,
	)
	if err != nil {
		return user, err
	}

	err = s.db.Model(&user).Updates(
		usermodel.User{
			Email:    newUser.Email,
			Password: hash,
		},
	).Error

	return user, err
}

func (s *Store) Delete(user usermodel.User) (usermodel.User, error) {
	user, err := s.GetByID(user.ID)
	if err != nil {
		return user, err
	}
	err := s.db.Delete(&user).Error
	return user, err
}
