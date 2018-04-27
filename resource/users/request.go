package users

import (
	"errors"
	"net/http"

	usermodel "github.com/moqafi/harper/model/user"
)

type UserRequest struct {
	*usermodel.User
}

// Bind validates required fields and values
func (ur *UserRequest) Bind(r *http.Request) error {
	if ur.Email == "" {
		return errors.New("email is required")
	}

	if ur.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
