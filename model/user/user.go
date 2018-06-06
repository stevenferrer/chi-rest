package user

import (
	model "github.com/moqafi/harper/model"
)

type User struct {
	model.Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password []byte `json:"password,omitempty" gorm:"not null"`
}

// Storer shall be implemented by database stores
type Storer interface {
	List(filter ...Filter) ([]User, error)
	GetByID(id uint64) (User, error)
	GetByEmail(email string) (User, error)

	// in create, update, and delete methods return user
	Create(User) (User, error)
	UpdateByID(id uint64, u User) (User, error)
	UpdateByEmail(email string, u User) (User, error)
	Delete(User) (User, error)
}

// Filter is used for filtering results
type Filter func(*FilterConfig) error

type FilterConfig struct {
	//User // easy method, inherit all User fields
	ID int64
}

func IDFilter(id int64) Filter {
	return func(fc *FilterConfig) error {
		fc.ID = id
		return nil
	}
}
