package main

import (
	"fmt"
	"sync"
)

var a sync.Mutex

var wg sync.WaitGroup

var wgrw sync.RWMutex

var syncmap sync.Map

func main() {
	a.Lock()
	fmt.Print("1, ")
	a.Unlock()
	fmt.Print("2, ")
	a.Unlock()
	fmt.Println("3")
}
