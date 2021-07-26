package gochannel

import (
	"fmt"
	"math/rand"
	"sync"
)

func player(name string, ch chan int) {
	defer wg.Done()
	for { // 因为不知道要打多少次球
		// 接球
		ball, ok := <-ch //怎样从通道里拿值, 返回两个值，当通道关闭，ok是false, 如果通道没值，虽然是死循环，但是也会卡在这里

		if !ok { // 通道关闭了，任务完成
			fmt.Printf("channel is closed! %s wins!", name)
			return // 退出for, 退出player
		}

		n := rand.Intn(100)

		if n%10 == 0 { // 有十分之一概率，把球打飞
			close(ch)
			fmt.Printf("%s misses the ball! %s loses\n", name, name)
			return
		}

		ball++
		fmt.Printf("%s receives ball %d\n", name, ball)

		ch <- ball // 把球传给对手
	}

}

var wg sync.WaitGroup

func Mainchannel() {

	wg.Add(2)
	ch := make(chan int) // unbuffered channel， 做为global var 或者传入goroutine中
	// ch <- 0, 非常有讲究，如果放在这，因为是没有缓冲区，没有人接，main会一直卡在这里，死锁，deadlock
	go player("heli", ch)
	go player("chong", ch)
	// 此时两个player会一直卡在channel取值处
	ch <- 0   // 启动两个goroutine
	wg.Wait() //直到两个player推出

}
