package users

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	// "github.com/go-chi/render"

	userModel "github.com/moqafi/harper/model/user"
)

func New() chi.Router {
	rs := usersResource{}
	return rs.routes()
}

type usersResource struct{}

// ctx middleware is used to load an user object from
// the URL parameters passed through as the request. In case
// the Article could not be found, we stop here and return a 404.
func (rs usersResource) ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var usr *userModel.User
		var err error

		if userID := chi.URLParam(r, "userId"); userID != "" && userID != "99" {
			// article, err = dbGetArticle(articleID)
			usr = &userModel.User{ID: userID}
		} else {
			err = errors.New("User not found")
		}
		// } else if articleSlug := chi.URLParam(r, "articleSlug"); articleSlug != "" {
		// 	article, err = dbGetArticleBySlug(articleSlug)
		// } else {
		// 	render.Render(w, r, ErrNotFound)
		// 	return
		// }

		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), "user", usr)
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

	r.Route("/{userId:[0-9]+}", func(r chi.Router) {
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
	w.Write([]byte("aaa list of users.."))
}

func (rs usersResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa create"))
}

func (rs usersResource) Get(w http.ResponseWriter, r *http.Request) {
	// Assume if we've reach this far, we can access the article
	// context because this handler is a child of the rs.ctx
	// middleware. The worst case, the recoverer middleware will save us.
	usr := r.Context().Value("user").(*userModel.User)

	// if err := render.Render(w, r, NewArticleResponse(article)); err != nil {
	// 	render.Render(w, r, ErrRender(err))
	// 	return
	// }

	w.Write([]byte("UserID is " + usr.ID))
}

func (rs usersResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa update"))
}

func (rs usersResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa delete"))
}
