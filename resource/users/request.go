package users

import (
	"encoding/json"
	"net/http"

	usermodel "github.com/sf9v/chi-rest/model/user"
)

type UserRequest struct {
	*usermodel.User
}

func (ur *UserRequest) UnmarshalJSON(data []byte) error {
	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	var u usermodel.User

	u.Email = v["email"].(string)
	u.Password = []byte(v["password"].(string))

	ur.User = &u

	return nil
}

// Bind validates required fields and values
func (ur *UserRequest) Bind(r *http.Request) error {
	err := ur.User.Validate()
	return err
}
