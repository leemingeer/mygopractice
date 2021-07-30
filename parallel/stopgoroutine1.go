package main

import "fmt"
import "time"

func do_stuff() int {
	fmt.Println("do_stuff is called")
	return 1
}

func main() {

	ch := make(chan int, 100)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case ch <- do_stuff():
			case <-done:
				fmt.Println("done received, closing channel...")
				close(ch)
				return
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()

	for i := range ch {
		fmt.Println("receive value: ", i)
	}

	fmt.Println("finish")
}