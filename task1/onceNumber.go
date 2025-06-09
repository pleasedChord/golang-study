package main

import "fmt"

/*
使用map的方式实现：
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func UseMap(numbers []int) {
	tempMap := make(map[int]int)
	for _, cli := range numbers {
		tempMap[cli] += 1
	}

	for k, v := range tempMap {
		if v == 1 {
			fmt.Println("只出现一次的数字是：", k)
		}
	}
}

/*
使用异或的方式实现
*/
func UseXOR(numbers []int) {
	result := 0
	for _, cli := range numbers {
		result ^= cli
	}
	fmt.Println("只出现一次的数字是：", result)
}

/*
只出现一次的数字
*/
func main() {
	numbers := []int{7, 2, 1, 2, 1}
	UseMap(numbers)
	UseXOR(numbers)
}
