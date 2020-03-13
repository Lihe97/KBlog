package controllers

import (
	"KBlog/models"
	"time"
)

type AddArticleController struct {
	//检查登陆状态？
	BaseController
}

func(this *AddArticleController) Get(){
	this.TplName = "write_article.html"
}
func(this *AddArticleController) Post(){
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")
	//author := this.GetString("author")

	article := models.Article{0,title,tags,short,content,"kansas",time.Now().Unix()}
	_,err := models.AddArticle(article)

	var response map[string]interface{}

	if err!=nil{
		response = map[string]interface{}{"code":0,"message":"写文章失败"}
	}else{
		response = map[string]interface{}{"code":1,"message":"写文章成功"}
	}
	this.Data["json"] = response
	this.ServeJSON()

}