package main

import (
	"bee06/db_mysql"
	_ "bee06/db_mysql"
	_ "bee06/routers"
	"github.com/astaxie/beego"
)

func main() {

	db_mysql.ConnectDB()
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")


	beego.Run()
}


