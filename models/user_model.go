package models

import (
	"KBlog/utils"
	"fmt"
)

type User struct {
	Id int
	Username string
	Password string
	Status int //0正常，1删除
	Createtime int64
}
//插入
func InsertUser(user User)(int64,error){
	return utils.ModifyDB("insert into users(username,password,status,createtime) values(?,?,?,?)",
		user.Username,user.Password,user.Status,user.Createtime)
}
func QueryUserWightCon(con string)int{
	sql := fmt.Sprintf("select id from users %s",con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id

}
//根据用户名查询id
func QueryUserWithUsername(username string)int{
	sql := fmt.Sprintf("where username='%s",username)
	return QueryUserWightCon(sql)
}

