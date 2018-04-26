package users

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	urender "github.com/unrolled/render"

	usermodel "github.com/moqafi/harper/model/user"
)

func New(store usermodel.Storer) chi.Router {
	u := usermodel.User{ID: 1, Email: "stevenf@moqafi.io"}
	store.Create(u)
	rs := usersResource{
		store: store,
		r:     urender.New(),
	}
	return rs.routes()
}

type usersResource struct {
	store usermodel.Storer
	r     *urender.Render
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
			usr, err = rs.store.Get(id)
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

		ctx := context.WithValue(r.Context(), "user", &usr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Routes creates a REST router for the users resource
func (rs usersResource) routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", rs.List) // GET /users - read a list of users

	r.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Authenticator)
		r.Post("/", rs.Create) // POST /users - create new user and persist it
	})

	r.Route("/{id:[0-9]+}", func(r chi.Router) {
		// load user if found to request context
		r.Use(rs.ctx)
		r.Get("/", rs.Get) // GET /users/{id} - read a single todo by :id

		r.Route("/", func(r chi.Router) {
			r.Use(jwtauth.Authenticator)
			r.Post("/", rs.Create)
			r.Put("/", rs.Update)    // PUT /users/{id} - update a single todo by :id
			r.Delete("/", rs.Delete) // DELETE /users/{id} - delete a single todo by :id
		})
	})

	return r
}

func (rs usersResource) List(w http.ResponseWriter, r *http.Request) {
	users, _ := rs.store.List()
	rs.r.JSON(w, http.StatusOK, users)
}

func (rs usersResource) Create(w http.ResponseWriter, r *http.Request) {
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
	render.Render(w, r, &UserResponse{Payload: user, Message: "User has been created"})
}

func (rs usersResource) Get(w http.ResponseWriter, r *http.Request) {
	// Assume if we've reach this far, we can access the article
	// context because this handler is a child of the rs.ctx
	// middleware. The worst case, the recoverer middleware will save us.
	usr := r.Context().Value("user").(*usermodel.User)

	rs.r.JSON(w, http.StatusOK, usr)
}

func (rs usersResource) Update(w http.ResponseWriter, r *http.Request) {
	usr := r.Context().Value("user").(*usermodel.User)

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
	render.Render(w, r, &UserResponse{Payload: user, Message: "User has been updated"})
}

func (rs usersResource) Delete(w http.ResponseWriter, r *http.Request) {
	usr := r.Context().Value("user").(*usermodel.User)

	err := rs.store.Delete(*usr)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, &UserResponse{Payload: usr, Message: "User has been deleted"})
}
