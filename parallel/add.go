package main

import (
	"fmt"
	"sync/atomic"
	"sync"
)

var count1 int32
var mu sync.Mutex
var wg sync.WaitGroup

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//mu.Lock()
			//count1 = count1 + 1
			//mu.Unlock()
			atomic.AddInt32(&count1, 1)
		}()
	}
	wg.Wait()
	//time.Sleep(time.Second)
	fmt.Println(count1)
}
