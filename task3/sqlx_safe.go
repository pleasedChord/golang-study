package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct {
	Id     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func initData() {
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("连接db失败")
	}

	//如果表已经存在则drop
	if db.Migrator().HasTable(&Book{}) {
		db.Migrator().DropTable(&Book{})
	}

	//建表
	db.AutoMigrate(&Book{})

	//初始化数据
	books := []Book{
		{Title: "世界历史", Author: "爱谁谁", Price: 99},
		{Title: "中国地理", Author: "我不知道啊", Price: 86},
		{Title: "舌尖上的中国", Author: "你说呢", Price: 32},
	}

	for _, book := range books {
		db.Create(&book)
	}
}

func main() {
	initData()

	db, err := sqlx.Open("sqlite3", "books.db")
	if err != nil {
		fmt.Println("连接db失败")
	}

	var books []Book
	query := `select id, title, author, price from books where price > ?`
	if err := db.Select(&books, query, "50"); err != nil {
		fmt.Println("select err:", err)
		return
	} else {
		for _, book := range books {
			fmt.Println(book)
		}
	}

}
