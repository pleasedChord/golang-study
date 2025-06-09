package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person     Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Println("员工id:", e.EmployeeID, "员工姓名:", e.Person.Name, "员工年龄:", e.Person.Age)
}

/*
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
*/
func main() {
	employee := Employee{
		Person: Person{
			Name: "张三",
			Age:  33,
		},
		EmployeeID: "12345",
	}

	employee.PrintInfo()
}
