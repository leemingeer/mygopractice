package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for i := 0; i < 7; i++ {
		fmt.Println("read from channel: ", <-ch)
	}
}
