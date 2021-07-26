package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func say(id string) {
	time.Sleep(time.Second)
	fmt.Println("I am done! id: " + id)
	wg.Done() // 任务完成
}

var wg sync.WaitGroup

func Maingoroutine() {
	// nothing output
	//go say("Hello")
	//go say("World")

	wg.Add(3) // 总共有3个任务

	go say("Hello")
	go say("World")
	// 匿名函数
	go func(id string) {
		// 立刻打印，等1s后，两个say()才会执行
		fmt.Println(id)
		wg.Done()
	}("here")

	wg.Wait()

}
