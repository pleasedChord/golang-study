package main

import (
	"fmt"
	"sync"
	"time"
)

/*
编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
*/
func base() {
	c := make(chan int)
	defer close(c)

	//生成从1到10的整数，并将这些整数发送到通道中
	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
		}
	}()

	//从通道中接收这些整数并打印出来
	go func() {
		for num := range c {
			fmt.Println("c===", num)
		}
	}()

	time.Sleep(2 * time.Second)
}

/*
实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，
消费者协程从通道中接收这些整数并打印。
*/
func bufferChannel(bufferNum int) {
	c := make(chan int, bufferNum)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 100; i++ {
			c <- i
		}

		close(c)
	}()

	go func() {
		for num := range c {
			fmt.Println("c==", num)
		}
	}()

	wg.Wait()

	time.Sleep(5 * time.Second)
}

func main() {
	//无缓冲的channel
	base()
	//有缓冲的channel
	bufferChannel(10)
}
