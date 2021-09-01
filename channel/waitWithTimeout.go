package main

import (
	"fmt"
	"time"
)

func main() {
	// add in goroutine
	i := 1
	ch := make(chan int)
	go func() {
		i += 1
		time.Sleep(3 * time.Second)
		ch <- 1
		close(ch) // 如果这里没有close, ch一直没有关闭，ch会卡住
		// 关闭ch后， 从关闭的ch里取值，res一直取到0值，是不阻塞的，所以应该加return返回

	}()

	for {
		select {
		case res := <-ch:
			fmt.Printf("finally: %v, %d\n", i, res)
			return

		case <-time.After(5 * time.Second):
			fmt.Println("timeout")
		}
	}
}