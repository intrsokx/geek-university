package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 用户信息
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	//  db, err := gorm.Open("mysql", "user:password@(localhost)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	//mysql 8.0.22版本的密码信息不能加主机信息（eg1）,必须使用eg2格式
	//eg1:   root:mysqlPass@123@(localhost:3316)@/test
	//eg2:   root:mysqlPass@123@/test
	db, err := gorm.Open("mysql", "root:mysqlPass@123@/test")
	if err != nil {
		fmt.Printf("create db err:%v", err)
		return
	}
	defer db.Close()

	db.AutoMigrate(&UserInfo{})

	u1 := UserInfo{3, "七米", "男", "篮球"}
	u2 := UserInfo{4, "沙河娜扎", "女", "足球"}
	// 创建记录
	db.Create(&u1)
	db.Create(&u2)
	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	db.Delete(&u)
}
