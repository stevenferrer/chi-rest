package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/sirupsen/logrus"

	"github.com/sf9v/chi-rest/middleware/logger"
	usersresource "github.com/sf9v/chi-rest/resource/users"
	userstore "github.com/sf9v/chi-rest/store/user/memory"
)

// TODO: Move this out of global scope
var tokenAuth *jwtauth.JWTAuth

func init() {
	var signKey = []byte("this is a secret")
	tokenAuth = jwtauth.New("HS256", signKey, nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{
		"userId": 1,
		"email":  "foo@bar.com",
	})
	log.Printf("DEBUG: a sample jwt is %s\n", tokenString)
}

// Run starts the application
func Run() {
	port := ":3333"
	host := "localhost"
	addr := host + port

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		log.Printf("Server running on %s\n", addr)
		//		db := getUserModelDB()
		//		defer db.Close()
		var db *gorm.DB
		errs <- http.ListenAndServe(addr, router(db))
	}()

	log.Fatalf("exit: %s", <-errs)
}

func router(db *gorm.DB) http.Handler {
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
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seekd, verify and validate JWT tokens
		// r.Use(jwtauth.Verifier(tokenAuth))

		// Note: jwtauth.Authenticator should be
		// added by different routes. For example,
		// some routes allow GET and disallow POST

		// Initialize Model Stores
		userStore := userstore.New()
		r.Mount("/users", usersresource.New(userStore, tokenAuth))
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			// Show API Banner or something
		})

		r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
			// Show API Docs
		})
	})

	return r
}
