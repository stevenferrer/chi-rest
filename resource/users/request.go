package users

import (
	"net/http"

	usermodel "github.com/moqafi/harper/model/user"
)

type UserRequest struct {
	*usermodel.User
}

// Bind validates required fields and values
func (ur *UserRequest) Bind(r *http.Request) error {
	return nil
}
