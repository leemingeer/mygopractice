
package main

import (
"fmt"
"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i // 因为capacity是2，写入0,1时阻塞
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second) // write goroutine执行
	for v := range ch {
		fmt.Println("read value", v,"from ch")
		time.Sleep(2 * time.Second) // 当从channel读出一个值时，write goroutine才会往里塞入一个数值，然后又阻塞了

	}
}