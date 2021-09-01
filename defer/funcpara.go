package main

import "fmt"

func main() {
	i := 0
	defer fmt.Println("a:", i)
	//闭包调用，将外部i传到闭包中进行计算，不会改变i的值，如上边的例3
	i++
	defer func(i int) {
		fmt.Println("b:", i)
	}(i)
	//闭包调用，捕获同作用域下的i进行计算
	defer func() {
		fmt.Println("c:", i)
	}()
	i++
}