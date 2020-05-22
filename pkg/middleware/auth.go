package middleware

import (
	"context"
	"golang-stepik-2020q1/5/99_hw/redditclone/pkg/ljwt"
	"golang-stepik-2020q1/5/99_hw/redditclone/pkg/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Middleware represents a wrapper extending the handler's functionality
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// AuthorizedUserMiddleware checks that a user is authorized to make a call
func AuthorizedUserMiddleware() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			tokenString := r.Header.Get("Authorization")
			if len(tokenString) == 0 {
				jsonMessage := utils.GetJSONMessageAsString("Missing Authorization Header")
				http.Error(w, jsonMessage, http.StatusUnauthorized)
				return
			}
			tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

			claims := jwt.MapClaims{}

			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(ljwt.ServerKey), nil
			})

			if err != nil || !token.Valid {
				jsonMessage := utils.GetJSONMessageAsString(err.Error())
				http.Error(w, jsonMessage, http.StatusUnauthorized)
				return
			}

			var interf jwt.MapClaims
			var tokenUsername string
			var tokenUserID string
			var ok bool
			if interf, ok = claims["user"].(map[string]interface{}); !ok {
				jsonMessage := utils.GetJSONMessageAsString("Error getting user's authorisations")
				http.Error(w, jsonMessage, http.StatusUnauthorized)
				return
			}
			if tokenUsername, ok = interf["username"].(string); !ok {
				jsonMessage := utils.GetJSONMessageAsString("Error getting user's authorisation attribute")
				http.Error(w, jsonMessage, http.StatusUnauthorized)
				return
			}
			if tokenUserID, ok = interf["id"].(string); !ok {
				jsonMessage := utils.GetJSONMessageAsString("Error getting user's authorisation attribute")
				http.Error(w, jsonMessage, http.StatusUnauthorized)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx,
				utils.CurrentUserKey,
				&ljwt.TokenUser{
					Username: tokenUsername,
					ID:       tokenUserID,
				})

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}

}
