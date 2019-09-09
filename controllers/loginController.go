package controllers

import (
	"github.com/kataras/iris"
	"os"
	"iris-init/common"
		"github.com/dgrijalva/jwt-go"
	"time"
	"iris-init/models"
		"log"
	"iris-init/service"
	)

func Login(ctx iris.Context)  {
	name := os.Getenv("DB_DIALECT")
	ctx.WriteString(name)
}

func AccountLogin(ctx iris.Context) {
	ctx.Application()
	loginAccount := new(models.LoginAccount)
	if err := ctx.ReadJSON(&loginAccount); err != nil {
		ctx.JSON(common.ErrorClientParams)
		return
	}
	if err := checkQueryParams(loginAccount); err != nil {
		ctx.JSON(err)
		return
	}

	var userAccount models.UserAccount

	err := service.GetById(&userAccount, loginAccount.Account)

	if err != nil {
		ctx.JSON(common.ErrUserEmpty)
		return

	}
	//passwordByt, err := base64.StdEncoding.DecodeString(loginAccount.Password)
	//if err != nil {
	//	ctx.JSON(common.ErrUserOrPwd)
	//	return
	//}
	//
	//password := utils.HmacSHA1Base64Encrypt(string(passwordByt), userAccount.Openid)
	//fmt.Println(userAccount)
	//
	//fmt.Println(password)
	if loginAccount.Password != userAccount.Openid {
		ctx.JSON(common.ErrUserOrPwd)
		return
	}
	ctx.JSON(models.RLoginAccountInfo{
		Account: userAccount.Id,
		Token:   getAccessTokenString(userAccount.Id),
	})

}

func checkQueryParams(account *models.LoginAccount) (err error) {
	if account.Account < 99999 {
		//err = common.ErrUserEmpty
	}
	//TODO
	return err
}

func getAccessTokenString(userId int64) string {
	claims := make(jwt.MapClaims)
	claims["userId"] = userId
	claims["version"] =	os.Getenv("application.version")
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	tokenString, err := token.SignedString([]byte("nooneisperfectbutme.*#"))
	if err != nil {
		log.Println("getAccessTokenString err", err)
		return ""
	}
	return tokenString
}
