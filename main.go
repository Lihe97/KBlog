package main

import (
	_ "KBlog/routers"
	"KBlog/utils"
	"github.com/astaxie/beego"
)

func main() {
	utils.InitMysql()
	//2种方法配置session
	//beego.SessionConfig{}.SessionOn =
	
	beego.Run()



}

