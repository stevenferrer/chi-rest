package users

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	urender "github.com/unrolled/render"

	usermodel "github.com/moqafi/harper/model/user"
)

// TODO: make resource handlers private

type key string

const userKey key = "user"

func createDummyUsers() []usermodel.User {
	var users []usermodel.User
	for i := 0; i < 10; i++ {
		u := usermodel.User{
			Email:    fmt.Sprintf("sample%d@example.com", i+1),
			Password: fmt.Sprintf("sample%d", i+1),
		}
		users = append(users, u)
	}

	return users
}

func New(store usermodel.Storer, tokenAuth *jwtauth.JWTAuth) chi.Router {
	users := createDummyUsers()
	for _, u := range users {
		store.Create(u)
	}

	rs := usersResource{
		store:     store,
		r:         urender.New(),
		tokenAuth: tokenAuth,
	}
	return rs.routes()
}

type usersResource struct {
	store     usermodel.Storer
	tokenAuth *jwtauth.JWTAuth
	r         *urender.Render
}

// ctx middleware is used to load an user object from
// the URL parameters passed through as the request. In case
// the Article could not be found, we stop here and return a 404.
func (rs usersResource) ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var usr usermodel.User
		var err error

		if userID := chi.URLParam(r, "id"); userID != "" {
			id, _ := strconv.ParseInt(userID, 10, 64)
			// usr = &usermodel.User{ID: id}
			usr, err = rs.store.GetByID(id)
			if err != nil {
				render.Render(w, r, ErrInvalidRequest(err))
				return
			}
		} else {
			render.Render(w, r, ErrInvalidRequest(errors.New("User not found")))
			return
		}

		if err != nil {
			// http.Error(w, http.StatusText(404), 404)
			render.Render(w, r, ErrRender(err))
			return
		}

		ctx := context.WithValue(r.Context(), userKey, &usr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Routes creates a REST router for the users resource
func (rs *usersResource) routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", rs.list)      // GET /users - read a list of users
	r.Post("/auth", rs.auth) // jwt token auth

	r.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Authenticator)
		r.Post("/", rs.create) // POST /users - create new user and persist it
	})

	r.Route("/{id:[0-9]+}", func(r chi.Router) {
		// load user if found to request context
		r.Use(rs.ctx)
		r.Get("/", rs.get) // GET /users/{id} - read a single todo by :id

		r.Route("/", func(r chi.Router) {
			r.Use(jwtauth.Authenticator)
			r.Post("/", rs.create)
			r.Put("/", rs.update)    // PUT /users/{id} - update a single todo by :id
			r.Delete("/", rs.delete) // DELETE /users/{id} - delete a single todo by :id
		})
	})

	return r
}

func (rs *usersResource) list(w http.ResponseWriter, r *http.Request) {
	users, _ := rs.store.List()

	if err := render.RenderList(w, r, NewUserListResponse(users)); err != nil {
		render.Render(w, r, ErrRender(err))
	}
}

func (rs *usersResource) create(w http.ResponseWriter, r *http.Request) {
	data := &UserRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	user := data.User

	err := rs.store.Create(*user)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusCreated)
	render.Render(w, r, &UserResponse{User: user, Message: "User has been created"})
}

func (rs *usersResource) get(w http.ResponseWriter, r *http.Request) {
	// Assume if we've reach this far, we can access the user
	// context because this handler is a child of the rs.ctx
	// middleware. The worst case, the recoverer middleware will save us.
	usr := r.Context().Value(userKey).(*usermodel.User)

	if err := render.Render(w, r, NewUserResponse(*usr)); err != nil {
		render.Render(w, r, ErrRender(err))
	}
}

func (rs *usersResource) update(w http.ResponseWriter, r *http.Request) {
	usr := r.Context().Value(userKey).(*usermodel.User)

	data := &UserRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	user := data.User

	// set the original id
	user.ID = usr.ID

	err := rs.store.Update(*user)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusCreated)
	render.Render(w, r, &UserResponse{User: user, Message: "User has been updated"})
}

func (rs *usersResource) delete(w http.ResponseWriter, r *http.Request) {
	usr := r.Context().Value(userKey).(*usermodel.User)

	err := rs.store.Delete(*usr)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, &UserResponse{User: usr, Message: "User has been deleted"})
}

func (rs *usersResource) auth(w http.ResponseWriter, r *http.Request) {
	data := &UserRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	authUser := data.User

	user, err := rs.store.GetByEmail(authUser.Email)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	if user.Password != authUser.Password {
		render.Render(w, r, ErrInvalidRequest(errors.New("User authentication failed")))
		return
	}

	_, tokenString, err := rs.tokenAuth.Encode(jwtauth.Claims{
		"userId": user.ID,
		"email":  user.Email,
	})
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	rs.r.JSON(w, http.StatusOK, map[string]string{"token": tokenString})
}
