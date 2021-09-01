package main

import "sync"

var wg sync.WaitGroup
func main(){
	//ch := make(chan int)
	//<- ch
	wg.Add(1)
	//go func() {
	//	fmt.Println("come into gourine")
	//	for {
	//		time.Sleep(360 * time.Second)
	//
	//	}
	//}()
	go func(){

		wg.Done()
	}()
	wg.Wait()

}