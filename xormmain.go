package main

import (
	"fmt"
	"time"
	"xorm.io/xorm"
	_ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	Id int64
	Name string
	Salt string
	Age int
	Passwd string `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
	UserName string `xorm:"varchar(255)"`
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/cis?charset=utf8")
	if err != nil {
		fmt.Print(err)
		return
	}

	//engine.Sync(new(UserInfo))
	//results, err := engine.Query("select * from user")
	
	//user := &UserInfo{}
	userList := make([]UserInfo, 0)
	err = engine.Where("name = ?", "miguel").Find(&userList)
	fmt.Println(err)

}
