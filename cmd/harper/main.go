package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/sirupsen/logrus"

	"github.com/moqafi/harper/middleware/logger"
	usersResource "github.com/moqafi/harper/resource/users"
)

// TODO: Move this out of global scope
var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(jwtauth.Claims{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func main() {
	port := ":3333"
	host := "localhost"
	addr := host + port
	log.Printf("Server running on %s\n\n", addr)
	err := http.ListenAndServe(addr, router())
	if err != nil {
		log.Fatal(err)
	}
}

func router() http.Handler {
	r := chi.NewRouter()

	// Setup the logger backend using sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/sirupsen/logrus
	loggerLogrus := logrus.New()
	loggerLogrus.Formatter = &logrus.JSONFormatter{
		// disable, as we set our own
		DisableTimestamp: true,
	}

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(logger.NewStructuredLogger(loggerLogrus))
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seekd, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Note: jwtauth.Authentiator should be
		// added by different routes. For example,
		// some routes allow GET and disallow POST

		r.Mount("/users", usersResource.New())
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome Anonymous!"))
		})
	})

	return r
}
