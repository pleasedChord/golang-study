package main

import (
	"fmt"
	"time"
)

// 打印1-10奇数
func Rountine1() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

// 打印1-10偶数
func Rountine2() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	go Rountine1()
	go Rountine2()

	time.Sleep(2 * time.Second)
}
