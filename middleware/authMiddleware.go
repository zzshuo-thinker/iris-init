package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"fmt"
)


// 判断token是否有效
func AuthMiddleware(ctx iris.Context)  {
	token := ctx.Values().Get("jwt").(*jwt.Token)
	fmt.Println("token: ", token)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//ctx.Values().Set("token", int64(claims["token"].(float64)))
		ctx.Values().Set("uid", claims["userId"])
	} else {
		return
	}
	ctx.Next()
}
