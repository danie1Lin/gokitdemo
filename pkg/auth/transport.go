package auth

import (
	"github.com/go-kit/kit/auth/jwt"

	"context"
	"net/http"
)

func AddAuthorizationToJWTTokenContext(ctx context.Context, request *http.Request) context.Context {
	ctx = context.WithValue(ctx, jwt.JWTTokenContextKey, request.Header.Get("Authorization"))
	return ctx
}
