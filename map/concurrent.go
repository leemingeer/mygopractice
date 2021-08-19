package main

import (
	"fmt"
	"sync"
	"time"
)


var (
	d = make(map[string]string)
	rwm sync.RWMutex
)

func get(k string ) string{
	rwm.RLock()
	defer rwm.RUnlock()
	return d[k]
}

func set(k, v string ){
	rwm.Lock()
	defer rwm.Unlock()
	d[k] = v // 引用传递
}

func main(){
	go set("a", "A")
	go set("b", "B")
	time.Sleep(1 * time.Second)
	fmt.Println(d)
	go get("a")
	go get("b")
}

