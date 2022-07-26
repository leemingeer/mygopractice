package main

import (
	"fmt"
)

func worker2(ch1, ch2 <-chan int, stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		case job1 := <-ch1:
			fmt.Println(job1)
		case job2 := <-ch2:
		priority:
			for {
				select {
				case job1 := <-ch1:
					fmt.Println(job1)
				default:
					break priority
				}
			}
			fmt.Println(job2)
		}
	}
}
func main(){
	ch1, ch2 := make(chan int), make(chan int)
	stopch := make(chan struct{})

	worker2(ch1, ch2, stopch)

	ch1 <- 1
	ch2 <- 2
}