package middleware

import (
	"context"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/ljwt"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/session"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
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

// AuthorizedUserMiddleware makes sure that a user is authorized to make a call
func AuthorizedUserMiddleware(sm *session.SessionsManager, logger *zap.SugaredLogger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			tokenString := r.Header.Get("Authorization")
			logger.Infow("Auth middleware",
				"token", tokenString,
			)
			if len(tokenString) < len("Bearer ") {
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
			var sessionID string
			var tokenUserID string
			var ok bool
			if interf, ok = claims["user"].(map[string]interface{}); !ok {
				jsonMessage := utils.GetJSONMessageAsString("Error getting user's authorisations")
				http.Error(w, jsonMessage, http.StatusUnauthorized)
				return
			}
			if sessionID, ok = claims["sessionId"].(string); !ok {
				jsonMessage := utils.GetJSONMessageAsString("Error getting Session Id")
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

			logger.Infow("Auth passed",
				"userId", tokenUserID,
				"userName", tokenUsername,
				"sessionId", sessionID,
			)

			sess, err := sm.Check(sessionID)
			if err != nil {
				logger.Infow("Error when checking session",
					"token_string", tokenString,
				)
				http.Error(w, `Unauthorized`, http.StatusUnauthorized)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, session.SessionKey, sess)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
