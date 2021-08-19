package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main(){

	for i:=0;i< 10; i++{
		wg.Add(1)
		go func(){
			fmt.Println(i)
			wg.Done()
		}() // 如果不传入i作为入参，i取的是闭包变量，外面i已经循环到10了，routine可能还没有执行

	}
	// 如果把这句放到for循环中，保证goroutine的顺利执行，也能达到顺序输出的效果
	// time.Sleep(1 * time.Second)
	wg.Wait()
}
