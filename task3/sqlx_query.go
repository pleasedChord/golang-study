package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Employee struct {
	Id         int
	Name       string
	Department string
	Salary     float64
}

func initData() {
	db, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("连接db失败")
	}

	//如果表已经存在则drop
	if db.Migrator().HasTable(&Employee{}) {
		db.Migrator().DropTable(&Employee{})
	}

	//建表
	db.AutoMigrate(&Employee{})
	//舒适化数据
	employees := []Employee{
		{Name: "张三", Department: "技术部", Salary: 2000},
		{Name: "李四", Department: "人力资源部", Salary: 3500},
		{Name: "王五", Department: "技术部", Salary: 5000}}

	for _, emp := range employees {
		db.Create(&emp)
	}
}

func QueryEmpsByDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	var emps []Employee
	query := `select id, name, department, salary 
	from employees 
	where department  = ?`
	err := db.Select(&emps, query, department)
	return emps, err
}

func QueryTopSalaryEmp(db *sqlx.DB) (Employee, error) {
	var emp Employee
	query := `select * from employees order by salary DESC limit 1`
	err := db.Get(&emp, query)
	if err == sql.ErrNoRows {
		return Employee{}, err
	}

	return emp, err
}

func main() {
	//1.初始化数据
	//initData()

	//2.使用sqlx连接数据库
	db, err := sqlx.Open("sqlite3", "employees.db")
	if err != nil {
		fmt.Println("连接数据库失败:", err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("ping失败:", err)
	}

	//3.查询所有部门为 "技术部" 的员工信息
	emps, err := QueryEmpsByDepartment(db, "技术部")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("技术部员工:")
	for _, emp := range emps {
		fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n",
			emp.Id, emp.Name, emp.Department, emp.Salary)
	}
	fmt.Println()

	//4.查询工资最高的员工信息
	emp, err := QueryTopSalaryEmp(db)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("薪水最高的员工:")
	fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n",
		emp.Id, emp.Name, emp.Department, emp.Salary)
	fmt.Println()

}
