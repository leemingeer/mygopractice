package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan int)
	go func(){
		for {
			v, ok := <-ch
			if ok {
				fmt.Println(v)
			} else{
				fmt.Println("channel closed,stop read zero value from channel")
				break
			}
		}
	}()
	ch <- 2
	close(ch)
	time.Sleep(10*time.Second)
}
