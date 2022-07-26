package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

func calcAndStoreStats(){
	fmt.Println("Hello!")
}

func main(){
	stopChannel := make(chan struct{})
	go wait.JitterUntil(func() {
		calcAndStoreStats()
	}, 0*time.Second, 0, true, stopChannel)
    time.Sleep(3*time.Second)
}
