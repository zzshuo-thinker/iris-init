package models

import "time"

type UserAccount struct {
	Id    int64 `xorm:"pk"`
	Unionid     string
	Openid   	string
	Avatar      string
	Nickname    string
	CreatedAn time.Time
}

type LoginAccount struct {
	Account  int64  `json:"account"`
	Password string `json:"password"`
}

type RLoginAccountInfo struct {
	Account int64  `json:"account"`
	Token   string `json:"token"`
}

