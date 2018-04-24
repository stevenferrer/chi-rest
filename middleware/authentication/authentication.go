package authentication

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

//token *jwtauth.JWTAuth
func JWTAuthenticator(token *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	h := func(next http.Handler) http.handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

		}
		return http.HandlerFunc(fn)
	}

	return h
}
