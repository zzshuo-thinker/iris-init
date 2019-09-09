package routers

import (
	"github.com/kataras/iris"
	"iris-init/controllers"
			"net/http"
	"github.com/kataras/iris/core/router"
	"iris-init/middleware"
)

func Router(app *iris.Application)  {

	v1 := app.Party("/v1")
	{
		v1.Get("/login", controllers.AccountLogin)
		v1.PartyFunc("/user", func(user router.Party) {
			user.Use(middleware.JwtHandler().Serve, middleware.AuthMiddleware)
			user.Get("/info", func(ctx iris.Context) {
				ctx.WriteString("this is a text router")
			})
		})
	}



	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"code": http.StatusNotFound,
			"msg": "Not Found",
			"data": iris.Map{},
		})
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"code": http.StatusInternalServerError,
			"msg": "Internal Server Error",
			"data": iris.Map{},
		})
	})
}
