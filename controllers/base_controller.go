package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin bool
	//已经登陆用户的信息
	Loginuser interface{}
}

//判断是否已经登陆 重写prepare
func(this *BaseController)Prepare(){

	loginuser := this.GetSession("loginuser")
	fmt.Println("loginuser--------",loginuser)
	if loginuser != nil{
		this.IsLogin = true
		this.Loginuser = loginuser
	}else{
		this.IsLogin = false
	}
	this.Data["IsLogin"] = this.IsLogin
}