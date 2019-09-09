package config

import (
	"github.com/jinzhu/gorm"
	"os"
	)

var OrmConn *gorm.DB

func DatabaseConn() *gorm.DB {
	// 判断连接是否存在
	//if ormConn != nil || ormConn.DB().Ping() == nil {
	//	return ormConn
	//}
	conn, err := gorm.Open(os.Getenv("DB_DIALECT"), os.Getenv("DB_CONNECTION"))
	if err != nil {
		panic("connection error")
	}
	conn.DB().SetMaxOpenConns(100)
	conn.DB().SetMaxIdleConns(100)
	conn.DB().SetConnMaxLifetime(1000)
	conn.LogMode(true)

	OrmConn = conn

	return OrmConn
}

