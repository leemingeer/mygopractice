package main

import (
	"fmt"
	"sync"
	"time"
)

var mapInstance sync.Map
var mapMutex sync.Mutex
var key string

func worker(name string) {
	for {
		if _, exist := mapInstance.Load(key); exist {
			time.Sleep(1*time.Second)
			//fmt.Printf("%s drop\r\n", name)
			continue
		}
		time.Sleep(1*time.Second)
		mapInstance.Store(key, key)

		//do something
		fmt.Printf("%s work\r\n", name)
		time.Sleep(1*time.Second)
		mapInstance.Delete(key)
		fmt.Printf("%s drop\r\n", name)
	}

}


func worker1(name string) {
	for {
		mapMutex.Lock()
		if _, exist := mapInstance.Load(key); exist {
			mapMutex.Unlock()
			time.Sleep(1*time.Second)
			continue
		}

		time.Sleep(1*time.Second)
		mapInstance.Store(key, key)
		mapMutex.Unlock()

		//do something
		fmt.Printf("%s work\r\n", name)
		time.Sleep(1*time.Second)
		mapMutex.Lock()
		mapInstance.Delete(key)
		mapMutex.Unlock()
		fmt.Printf("%s drop\r\n", name)
	}

}

func main() {
	key = "mutex"
	for loop := 0; loop < 10; loop++ {
		go worker(fmt.Sprintf("student%d", loop))
	}
	time.Sleep(5*time.Minute)
}