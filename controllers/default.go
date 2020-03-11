package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "干死鲁能"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
