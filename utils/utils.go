package utils

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"log"
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
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}