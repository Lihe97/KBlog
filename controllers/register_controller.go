package controllers

import (
	"KBlog/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct{
	beego.Controller
}
func(this *RegisterController) Get(){
	this.TplName = "register.html"
}
func(this *RegisterController) Post(){
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	fmt.Println(username,password,repassword)

	//判断用户名重？能不能注册
	id := models.QueryUserWithUsername(username)
	fmt.Println("id:" ,id)
	if id >0{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"用户名已存在"}
		this.ServeJSON()
		return
	}

	//
	//password = utils.

}

