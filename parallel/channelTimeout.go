package main
import (
	"fmt"
	"time"
)
func main() {
	ch := make(chan int)
	quit := make(chan bool)
	//新开一个协程
	go func() {
		for { // keep select listening, other a matched case will cause select over
			select { // 只要其中有一个 case 已经完成，程序就会继续往下执行，而不会考虑其他 case 的情况
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true
			}
			fmt.Println("select goroutine over")
		}
	}() //别忘了()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	<-quit
	fmt.Println("程序结束")
}