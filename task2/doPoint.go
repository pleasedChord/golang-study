package main

import "fmt"

/*
编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
*/
func Test1(i *int) {
	ii := *i
	ii += 10
	fmt.Println(ii)
}

/*
实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
*/
func Test2(slice *[]int) {
	s := *slice
	for i := 0; i < len(s); i++ {
		s[i] = s[i] * 2
	}
}

func main() {
	i := 99
	Test1(&i)

	slice := []int{1, 2, 3, 4, 5}
	Test2(&slice)
	fmt.Println(slice)
}
