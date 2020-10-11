package routers

import (
	"bee06/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user_register",&controllers.RegisterController{})

    beego.Router("/log_html",&controllers.RegisterController{})

    beego.Router("/user_log",&controllers.LogController{})
}
