package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// 模拟用户请求数量
	requestCount := 10
	fmt.Println("goroutine_num", runtime.NumGoroutine())
	ch := make(chan bool, 20)
	for i := 0; i < 3; i++ {
		x := i
		go Read(ch, x)
	}

	for i := 0; i < requestCount; i++ {
		wg.Add(1)
		ch <- true
	}

	wg.Wait()
}

func Read(ch chan bool, y int) {
	for _ = range ch {
		fmt.Printf("goroutine_num: %d, go func: %d\n", runtime.NumGoroutine(), y)
		wg.Done()
	}
}
