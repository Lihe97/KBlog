package controllers

import (
	"KBlog/models"
	"KBlog/utils"
	"fmt"
	"github.com/astaxie/beego"
	"time"
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
	password = utils.MD5(password)
	fmt.Println("MD5后",password)

	user := models.User{0,username,password,0,time.Now().Unix()}
	_ , err := models.InsertUser(user)
	if err!= nil{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"注册失败"}

	}else{
		this.Data["json"] = map[string]interface{}{"code":1,"message":"注册成功"}
	}
	this.ServeJSON()

}

