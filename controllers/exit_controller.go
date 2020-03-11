package controllers

import "github.com/astaxie/beego"

type ExitController struct {
	beego.Controller
}

func(this *ExitController)Get(){

	this.DelSession("loginuser")
	this.Redirect("/",302)
}