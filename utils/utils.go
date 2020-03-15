package utils

import (
	"bytes"
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
	"html/template"
	"log"
	"time"
)

var db *sql.DB

func InitMysql(){
	fmt.Println("!!!!!!!!!!!!!!!!!!")
	fmt.Println("初始化MYsql")
	driverName := beego.AppConfig.String("driverName")

	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port +")/" + dbname + "?charset=utf8"
	fmt.Println(dbConn)
	db1,err := sql.Open(driverName,dbConn)
	if err != nil{
		fmt.Println("创建失败",err.Error())
	}else{
		fmt.Println("链接成功？")
		db = db1
		CreateTableWithUser()
		CreateTableWithArticle()
	}
}
func ModifyDB(sql string, args ...interface{}) (int64,error){
	result ,err :=db.Exec(sql,args...)
	if err !=nil{
		log.Println("失败",err)
		return 0,err
	}
	count ,err := result.RowsAffected()
	if err != nil{
		log.Println("失败",err)
		return 0,err
	}
	return count ,nil

}

func CreateTableWithUser(){
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	ModifyDB(sql)
}
func CreateTableWithArticle() {
	sql := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	ModifyDB(sql)
}
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}
func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}
func SwitchTimeStampToData(timeStamp int64) string {
	t := time.Unix(timeStamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

func SwitchMarkdownToHtml(content string) template.HTML {

	markdown := blackfriday.MarkdownCommon([]byte(content))

	//获取到html文档
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))

	/**
	对document进程查询，选择器和css的语法一样
	第一个参数：i是查询到的第几个元素
	第二个参数：selection就是查询到的元素
	*/
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
		fmt.Println(selection.Html())
		fmt.Println("light:", string(light))
		fmt.Println("\n\n\n")
	})
	htmlString, _ := doc.Html()
	return template.HTML(htmlString)
}