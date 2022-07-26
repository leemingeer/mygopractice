package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var concurrentNum = 10
var appName = "pod"
func main(){

	for c := 0; c < concurrentNum; c++ {
		wg.Add(1)
		curappName := appName + strconv.Itoa(c)
		// interval 1s, otherwise some obj are not created for completely concurrency
		time.Sleep(time.Second)
		go func() {
			defer wg.Done()
			fmt.Println(curappName)
		}()
	}
	fmt.Println("first main gorouting!")
	wg.Wait()
	fmt.Println("first main goroutine over!")

	for c := 0; c < concurrentNum; c++ {
		wg.Add(1)
		curappName := appName + strconv.Itoa(c)
		go func() {
			defer wg.Done()
            fmt.Println(curappName)
		}()
	}
	fmt.Println("second main gorouting!")
	wg.Wait()
	fmt.Println("second main goroutine over!")
}

