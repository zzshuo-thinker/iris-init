package middleware

import (
	"github.com/dgrijalva/jwt-go"
	jwtMiddleware "github.com/iris-contrib/middleware/jwt"
)

/**
 *	验证jwt
 */
func JwtHandler() *jwtMiddleware.Middleware {
	return jwtMiddleware.New(jwtMiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("nooneisperfectbutme.*#"), nil
		},

		SigningMethod: jwt.SigningMethodHS256,
	})

}
