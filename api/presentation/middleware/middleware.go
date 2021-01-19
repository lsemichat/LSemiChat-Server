package middleware

import (
	"app/api/constants"
	"app/api/infrastructure/lcontext"
	"app/api/infrastructure/lsession"
	"app/api/presentation/response"
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// CORS
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		token, err := lsession.GetSession(r)
		if err != nil {
			response.Unauthorized(w, errors.Wrap(err, "failed to get session"), "failed to authorization")
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		ctx = lcontext.SetUserID(ctx, claims[constants.JWTUserIDClaimsKey].(string))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
