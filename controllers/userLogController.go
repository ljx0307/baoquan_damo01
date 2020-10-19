package controllers

import (
	"bee06/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LogController struct {
	beego.Controller
}

/**
 * get方法处理/log.html的页面请求
 */
func (l*LogController) Get(){
	l.TplName = "log.html"
}
func (l*LogController)Post(){
	var user models.User
	err :=l.ParseForm(&user)
	if err != nil{
		fmt.Println(err.Error())
		//l.TplName = "error.html"
		l.Ctx.WriteString("抱歉，用户信息解析失败！")
		return
	}
	u, err := user.QueryUser()
	if err != nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户登录失败，请重试！")
		return
	}
	l.Data["Phone"]=u.Phone
	l.TplName = "deposit.html"
}
