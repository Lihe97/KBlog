package models

import "KBlog/utils"

type Article struct {
	Id int
	Title string
	Tags string
	Short string //简介来一个
	Content string
	Author string
	Createtime int64
}
func AddArticle(article Article)(int64,error){

	i,err := utils.ModifyDB("insert into article(title,tags,short,content,author,createtime)values(?,?,?,?,?,?)",
		article.Title,article.Tags,article.Short,article.Content,article.Author,article.Createtime)
	return i,err
}

