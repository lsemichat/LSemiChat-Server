package middleware

import (
	"app/api/constants"
	"app/api/infrastructure/lcontext"
	"app/api/infrastructure/lsession"
	"app/api/llog"
	"app/api/presentation/response"
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: 実装
		llog.Debug("auth middleware")
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
		// TODO: claimのkeyをconstantsにしたい
		ctx = lcontext.SetUserID(ctx, claims[constants.JWTUserIDClaimsKey].(string))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func MethodNotFoundHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: handlerで許可されているメソッド？
		// allowedMethods :=
		// llog.Debug(r.)
		// TODO: rのメソッドを確認する
		// method := r.Method

		next.ServeHTTP(w, r)
	})
}
