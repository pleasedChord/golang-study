package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，
每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/
func myLock() {
	var wg sync.WaitGroup
	var myLock sync.Mutex

	num := 0
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for i := 0; i < 1000; i++ {
				myLock.Lock()
				num++
				myLock.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("lock.num==", num)
}

/*
使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，
每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/
func myNolock() {
	var wg sync.WaitGroup
	var num int64

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddInt64(&num, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("atomic.num==", num)
}

func main() {
	myLock()
	myNolock()
}
