package main

import (
	"fmt"
	"time"
)

func G() {
	defer func() {
		//goroutine外进行recover
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
		fmt.Println("c")
	}()
	//创建goroutine调用F函数
	go F()
	fmt.Println("after goroutine")
	time.Sleep(time.Second)
}

func F() {
	defer func() {
		fmt.Println("b")
	}()
	//goroutine内部抛出panic
	panic("a")
}

func main(){
	G()
}