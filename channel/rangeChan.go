package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		data, ok := <-ch
		if !ok {
			return
		}
		fmt.Println("read from channel:", data)
	}
}
