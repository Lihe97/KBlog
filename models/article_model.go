package models

import (
	"KBlog/utils"
	"fmt"
	"github.com/astaxie/beego"
	"log"
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
func UpdateArticle(article Article) (int64, error) {
	//数据库操作
	// update article set title = ?, tags = ?, short = ?, content = ?  where id = ?
	return utils.ModifyDB("update article set title=?,tags=?,short=?,content=? where id=?",
		article.Title, article.Tags, article.Short, article.Content, article.Id)
}
func DeleteArticle(artID int) (int64, error) {
	//1、删除文章
	i, err := deleteArticleWithArtId(artID)
	//2、计算文章总页码数
	SetArticleRowsNum()
	return i, err
}

// delete from article where id = artID
func deleteArticleWithArtId(artID int) (int64, error) {
	return utils.ModifyDB("delete from article where id=?", artID)
}
func QueryArticlesWithTag(tag string) ([]Article, error) {
	sql := " where tags like '%&" + tag + "&%'"
	sql += " or tags like '%&" + tag + "'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "'"
	fmt.Println(sql)

	return QueryArticlesWithCon(sql)
}

func QueryArticleWithParam(param string) []string {
	// select tags from article
	rows, err := utils.QueryDB(fmt.Sprintf("select %s from article", param))
	if err != nil {
		log.Println(err)
	}
	var paramList []string
	for rows.Next() {
		arg := ""
		rows.Scan(&arg)
		paramList = append(paramList, arg)
	}
	return paramList
}