package main

import (
	"fmt"
	"sync"
)

func main() {

	s := "Hello World!A"
	b := []byte(s)
	ch := make(chan byte, len(s))
	for i := 0; i < len(s); i++ {
		ch <- b[i]
	}
	close(ch)

	var wg sync.WaitGroup
	start := make(chan bool)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			_, ok := <-start
			if ok {
				c, ok := <-ch
				if ok {
					fmt.Println("g1:", string(c))
				} else {
					close(start)
					return
				}
				start <- true
			} else {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_, ok := <-start
			if ok {
				c, ok := <-ch
				if ok {
					fmt.Println("g2:", string(c))
				} else {
					close(start)
					return
				}
				start <- true
			} else {
				return
			}
		}
	}()
	start <- true
	wg.Wait()
}
