package main

import (
	"fmt"
	"github.com/intrsokx/geek-university/week02/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	var db *gorm.DB
	var err error
	//dbName := "test"
	//dbName = ""

	//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	db, err = gorm.Open("mysql", fmt.Sprintf("root:123456@tcp(127.0.0.1:3306)"))
	if err != nil {
		fmt.Printf("create db err:%v", err)
		return
	}
	defer db.Close()
	//创建数据库
	//db.Exec("create database test")

	//切换数据库，并创建表
	db.Exec("use test").AutoMigrate(&entity.User{})
	//TODO 做一些初始化操作

	//db.CreateTable(entity.User{
	//	NO:     0,
	//	Name:   "康熙大哥哥",
	//	Gender: "男",
	//	Hobby:  "摸鱼",
	//})


	//db.CreateTable(entity.User{})
	//u := entity.User{
	//	NO:     0,
	//	Name:   "康熙大哥哥",
	//	Gender: "男",
	//	Hobby:  "摸鱼",
	//}
	//db.Create(u)
	//
	////插入10条记录
	//hobbys := []string{"打架", "泡妞", "抽烟", "篮球", "网球", "羽毛球"}
	//for i := 1; i < 11; i++ {
	//	uu := entity.User{
	//		NO:     uint(i),
	//		Name:   fmt.Sprintf("%s的第%d个小弟", u.Name, i),
	//		Gender: "男",
	//		Hobby:  hobbys[i%len(hobbys)],
	//	}
	//	db.Create(uu)
	//}

	//u1 := UserInfo{3, "七米", "男", "篮球"}
	//u2 := UserInfo{4, "沙河娜扎", "女", "足球"}
	//// 创建记录
	//db.Create(&u1)
	//db.Create(&u2)
	//// 查询
	//var u = new(UserInfo)
	//db.First(u)
	//fmt.Printf("%#v\n", u)
	//
	//var uu UserInfo
	//db.Find(&uu, "hobby=?", "足球")
	//fmt.Printf("%#v\n", uu)
	//
	//// 更新
	//db.Model(&u).Update("hobby", "双色球")
	//// 删除
	//db.Delete(&u)
}
