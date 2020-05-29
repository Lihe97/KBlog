package main

import (
	_ "KBlog/routers"
	"KBlog/utils"
	"github.com/astaxie/beego"

)

func main() {

	utils.InitMysql()
	go canalRun()

	beego.Run()


}


