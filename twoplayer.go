package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg 用来等待程序结束
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// main 是所有Go 程序的入口
func main() {
	// 创建一个无缓冲的通道
	court := make(chan int)
	// 计数加 2，表示要等待两个goroutine
	wg.Add(2)
	// 启动两个选手
	go player("Nadal", court)
	go player("Djokovic", court)
	// 发球
	fmt.Println("kick the ball!")
	court <- 1
	// 等待游戏结束
	wg.Wait()
}

// player 模拟一个选手在打网球
func player(name string, court chan int) {
	// 在函数退出时调用Done 来通知main 函数工作已经完成
	defer wg.Done()
	for {
		// 等待球被击打过来
		fmt.Printf("%s waiting ball coming!\n", name)
		ball, ok := <-court // get zero if channel already closed
		fmt.Printf("%s get the ball!\n", name)
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("I'm %d, channel is closed!\n", ball)
			fmt.Printf("Player %s Won\n", name)
			for i := 0; i < 5; i++ {
				ball1, ok1 := <-court
				fmt.Printf("channel is closed, but you can continue read ball: %d, ok2:%t from the closed channel\n", ball1, ok1)
			}
			fmt.Println("over")
			return
		}
		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		fmt.Printf("n: %d\n", n)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			// 关闭通道，表示我们输了
			fmt.Println("Close the channel!\n")
			close(court)
			return
		}
		// 显示击球数，并将击球数加1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		// 将球打向对手
		// you will get panic if channel is closed
		court <- ball
	}

	fmt.Printf("%s is over!", name)
}
