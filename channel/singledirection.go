package main

import (
	"fmt"
	"time"
)

func producer(chnl chan int) {
	for i := 0; i < 2; i++ {
		fmt.Printf("producing %d\n", i)
		chnl <- i
		fmt.Printf("produce %d ok\n ", i)
	}
	close(chnl)
	fmt.Println("producer has been closed!")
}
func main() {
	ch := make(chan int)
	go producer(ch)
	for {

		v, ok := <-ch
		if ok == false {
			fmt.Println("channel is closed!")
			break
		}
		fmt.Println("Received ", v, ok)
		time.Sleep(10 * time.Second)
	}
}