package common

import "fmt"

type Err struct {
	Code int 	`json:"code"`
	Msg string	`json:"msg"`
}

func (err *Err) Error() string {
	return fmt.Sprintf("ErrorCode:%d, ErrorMsg: %s", err.Code, err.Msg)
}

func (err *Err) SetError(code int, msg string) Err {
	return Err{code, msg}
}

func (err *Err) GetError() string {
	return fmt.Sprintf("ErrorCode:%d, ErrorMsg: %s", err.Code, err.Msg)
}

var (
	Success 			= Err{Code: 200,   Msg: "SUCCESS"}
	ErrorClientParams	= Err{Code: 10400, Msg: "缺少参数"}
	ErrorUnKnow			= Err{Code: 10401, Msg: "未知错误"}
)

var (
	ErrUserEmpty = Err{Code: 100101, Msg: "用户不存在。"}
	ErrPwdEmpty  = Err{Code: 100102, Msg: "密码不能为空。"}
	ErrUserOrPwd = Err{Code: 100103, Msg: "帐号或密码错误。"}
)
