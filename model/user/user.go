package user

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	is "github.com/go-ozzo/ozzo-validation/is"

	model "github.com/moqafi/harper/model"
)

type User struct {
	model.Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password []byte `json:"password,omitempty" gorm:"not null"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.By(minPwdLen)),
	)
}

// minPwdLen validation for minimum password length
func minPwdLen(value interface{}) error {
	pwd, _ := value.([]byte)
	if len(pwd) < 8 {
		return errors.New("should be at least 8 characters")
	}
	return nil
}
