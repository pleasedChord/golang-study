package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Grade string
}

func NewStudent(name string, age int, grade string) *Student {
	student := &Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
	return student
}

func InsertStd(db *gorm.DB, name string, age int, grade string) {
	std := NewStudent(name, age, grade)
	db.Create(std)
}

func main() {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("连接db失败")
	}
	db.AutoMigrate(&Student{})

	var std Student
	var stds []Student

	//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"
	// InsertStd(db, "张三", 20, "三年级")
	// InsertStd(db, "李四", 10, "一年级")
	// InsertStd(db, "王五", 30, "五年级")

	QueryAll(db)

	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	db.Where("age > ?", 18).Find(&stds)
	for _, cli := range stds {
		fmt.Println("查询年龄大于18的学生:", cli)
	}

	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	result := db.Model(&std).Where("name = ?", "张三").Update("grade", "四年级")
	if result.Error != nil {
		fmt.Println("更新失败")
	} else if result.RowsAffected == 0 {
		fmt.Println("没有更新任何数据")
	} else {
		fmt.Println("更新成功")
		QueryAll(db)
	}

	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	result = db.Delete(&std, "age < ?", "15")
	if result.Error != nil {
		fmt.Println("删除失败")
	} else if result.RowsAffected == 0 {
		fmt.Println("没有删除任何数据")
	} else {
		fmt.Println("删除成功")
		QueryAll(db)
	}

	//还原数据
	ReSet(db)
}

func QueryAll(db *gorm.DB) {
	var stds []Student
	db.Find(&stds)
	for _, cli := range stds {
		fmt.Println(cli)
	}
}

func ReSet(db *gorm.DB) {
	var std Student
	InsertStd(db, "李四", 10, "一年级")
	db.Model(&std).Where("name = ?", "张三").Update("grade", "三年级")
	db.update(&std, "")
}
