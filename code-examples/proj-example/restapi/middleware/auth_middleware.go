package middleware

import (
	"fmt"
	"net/http"
	"server/core"
	"server/domain"
)

type AuthMiddleware interface {
	CheckAuth(next http.Handler) http.Handler
	CheckRole(next http.Handler) http.Handler
}

type authMiddleware struct {
	authGuard core.AuthGuard
}

func NewAuthMiddleware(authGuard core.AuthGuard) AuthMiddleware {
	return &authMiddleware{
		authGuard: authGuard,
	}
}

func (middleware *authMiddleware) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("check auth for: ", r.URL.Path)

		login := ""

		if !middleware.authGuard.IsLoggedIn(login) {
			fmt.Println("Access denied")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (middleware *authMiddleware) CheckRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("check admin role for: ", r.URL.Path)

		login := "login"

		if !middleware.authGuard.IsAuthorized(login, domain.Admin) {
			fmt.Println("Access denied")
			return
		}

		next.ServeHTTP(w, r)
	})
}
