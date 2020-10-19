package routers

import (
	"bee06/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user_register",&controllers.RegisterController{})
	//http://127.0.0.1:8080/log.html
    //beego.Router("/log.html",&controllers.LogController{})
	//用户登录动作时的表单数据提交请求
    beego.Router("/user_log",&controllers.LogController{})
    beego.Router("/deposit.html",&controllers.LogController{})
    beego.Router("/upload",&controllers.FileUploadController{})
}
