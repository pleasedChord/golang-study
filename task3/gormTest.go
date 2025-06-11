package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserTemp struct {
	ID   uint
	Name string
	Age  int
}

func main() {
	// 连接 SQLite，数据库文件为当前目录下的 test.db
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("连接 SQLite 失败:", err)
	}

	// 自动创建表
	db.AutoMigrate(&User{})

	// 增删改查操作
	user := UserTemp{Name: "Alice", Age: 30}
	db.Create(&user)

	var result User
	db.First(&result, 1)
	log.Println("查询结果:", result)
}
