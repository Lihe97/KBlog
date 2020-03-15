package models

import (
	"KBlog/utils"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)
var articleRowsNum = 0


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
func FindArticleWithPage(page int) ([]Article, error) {
	//从配置文件中获取每页的文章数量
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	// 2页面， 每页显示2条 {3,4}
	// 3页面，每页显示5条  {11,12,13,14,15}
	fmt.Println("---------->page", page)
	return QueryArticleWithPage(page, num)
}

func  QueryArticlesWithTag(tag string) ([]Article, error) {
	sql := " where tags like '%&" + tag + "&%'"
	sql += " or tags like '%&" + tag + "'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "'"
	fmt.Println(sql)
	//sql: like

	// tags: http&web&socket&互联网&计算机
	//       http&web
	//       web&socket&互联网&计算机
	//       web

	// http://localhost:8080?tag=web

	// %&web&%   %代表任何内容都可以匹配
	// %&web
	// web&%
	// web
	return QueryArticlesWithCon(sql)
}
func GetArticleRowsNum() int {
	if articleRowsNum == 0 {
		articleRowsNum = QueryArticleRowNum()
	}
	return articleRowsNum
}

//查询文章的总条数
func QueryArticleRowNum() int {
	row := utils.QueryRowDB("select count(id) from article")
	num := 0
	row.Scan(&num)
	return num
}

//设置页数
func SetArticleRowsNum(){
	articleRowsNum = QueryArticleRowNum()
}
func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryArticlesWithCon(sql)
}

func QueryArticlesWithCon(sql string) ([]Article, error) {
	sql = "select id,title,tags,short,content,author,createtime from article " + sql
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content, author, createtime}
		artList = append(artList, art)
	}
	return artList, nil
}
func QueryArticleWithId(id int) Article {
	row := utils.QueryRowDB("select id,title,tags,short,content,author,createtime from article where id=" + strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{id, title, tags, short, content, author, createtime}
	return art
}