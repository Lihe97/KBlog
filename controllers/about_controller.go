package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {

	c.Data["author"] = "kansas"
	c.Data["tel"] = "Tel：153-1870-8097"
	c.Data["qq"] = "QQ：285270417"
	c.Data["github"] = "github: https://github.com/Lihe97"
	c.Data["CSDN"] = "csdn: https://blog.csdn.net/kansas_lh"


	c.TplName = "aboutme.html"
}
