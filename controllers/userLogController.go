package controllers

import (
	"bee06/models"
	"beego"
)

type LogController struct {
	beego.Controller
}
func (l*LogController) Get(){
	l.TplName = "log.html"
}
func (l*LogController)Post(){
	var user models.User
	err :=l.ParseForm(&user)
	if err != nil{
		//l.TplName = "error.html"
		l.Ctx.WriteString("抱歉，用户信息解析失败！")
		return
	}
	u, err := user.QueryUser()
	if err != nil {
		l.Ctx.WriteString("抱歉，用户登录失败，请重试！")
		return
	}
	l.Data["Phone"]=u.Phone
	l.TplName = "log.html"
}
