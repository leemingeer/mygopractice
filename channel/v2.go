package main

import (
"fmt"
)

func producerv2(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	fmt.Printf("begin close channel!")
	close(chnl)
	fmt.Printf("channel is closed!")
}
func main() {
	ch := make(chan int)
	go producerv2(ch)
	for v := range ch {
		fmt.Println("Received ",v)
	}
	fmt.Println("end channel!")
}