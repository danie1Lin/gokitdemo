package auth

import (
	stdjwt "github.com/golang-jwt/jwt/v4"

	"context"
	"fmt"
	"github.com/go-kit/kit/auth/casbin"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
)

func ParseJWTMiddleware() endpoint.Middleware {
	return jwt.NewParser(func(t *stdjwt.Token) (interface{}, error) {
		// 根據 jwt 的 header 找到對應的key
		// valid eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImlkIjoiMTIzIn0.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.VN65OhsyUff3VWfTKU03b63Kod0WoxyvZVB0savCvx4
		// invalid  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImlkIjoiMjIyIn0.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.gLx28XdOYSHoeWhS-FNQ55D681Hk_gj0g9oENlEooAg
		// TODO: 處理 auth 的 error code
		key := []byte{}
		if t.Header["id"] == "123" {
			key = []byte("1234")
		}
		return key, nil
	}, stdjwt.SigningMethodHS256, jwt.StandardClaimsFactory)
}

func RBACMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		ctx = context.WithValue(ctx, casbin.CasbinModelContextKey, "model.conf")
		ctx = context.WithValue(ctx, casbin.CasbinPolicyContextKey, "policy.conf")
		claim, ok := ctx.Value(jwt.JWTClaimsContextKey).(*stdjwt.StandardClaims)
		if !ok {
			return nil, fmt.Errorf("claims not found")
		}
		// try this eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImlkIjoiMTIzIn0.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.VN65OhsyUff3VWfTKU03b63Kod0WoxyvZVB0savCvx4
		// https://jwt.io/#debugger-io?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImlkIjoiMTIzIn0.eyJzdWIiOiIxMjM0NTY3OCIsIm5hbWUiOiJEYW5pZWwiLCJpYXQiOjE1MTYyMzkwMjJ9.vkNBjV7IYA7LLCrlPxcAWheoHVLSMvuncs7VnLFuRxw
		return casbin.NewEnforcer(claim.Subject, "event", "create")(next)(ctx, request)
	}
}
