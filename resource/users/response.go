package users

import (
	"net/http"

	"github.com/go-chi/render"

	usermodel "github.com/moqafi/harper/model/user"
)

// TODO: Create a Generic response struct that
//		will be composed by other resources.
//		Maybe look at other golang packages for ideas

type UserResponse struct {
	User    *usermodel.User `json:"user"`
	Message string          `json:"message,omitempty"`
}

func (ur *UserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// hide user password
	ur.User.Password = ""
	return nil
}

func NewUserResponse(user usermodel.User) *UserResponse {
	return &UserResponse{User: &user}
}

type UserListResponse []*UserResponse

func NewUserListResponse(users []usermodel.User) []render.Renderer {
	list := []render.Renderer{}
	for _, user := range users {
		list = append(list, NewUserResponse(user))
	}
	return list
}

//--
// Error response payloads & renderers
//--

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found"}
