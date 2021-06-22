//测试wait.Until() 的用途
package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

type stop struct {
}

func main() {
	stopCh := make(chan struct{})
	//初始化一个计数器
	i := 0
	go wait.Until(func() {
		fmt.Printf("----%d----\n", i)

		i++
	}, time.Second, stopCh)

	time.Sleep(time.Second * 10)

	stopCh <- stop{}
	// 下面的形式，也是可以的
	//stopCh < struct {
	//}{}
	// close(stopCh)  测试这个也可以

	fmt.Println("---上面的go routines 结束----")

	// 主程序，再休息3秒钟，再结束
	time.Sleep(time.Second * 3)

	fmt.Println("---主程序结束----")
}
