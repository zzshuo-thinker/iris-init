package main

import (
	_ "github.com/go-sql-driver/mysql"
	"iris-init/bootstrap"
)

func main()  {
	bootstrap.App().Run()
}
