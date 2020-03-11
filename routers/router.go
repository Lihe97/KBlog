package routers

import (
	"KBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {

    beego.Router("/", &controllers.MainController{})

    //注册
    beego.Router("/register",&controllers.RegisterController{})
    //登陆
    beego.Router("/login",&controllers.LoginController{})
}
