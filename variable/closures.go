package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++ // 没有通过函数参数传入，i往上层找到i:=0
		return i
	}
}

func main() {

	nextInt := intSeq() // 函数执行完返回另一个函数， 因为返回的函数可能随时被执行多次，所以其引用的闭包变量会逃逸到堆上，何时被回收：闭包函数对象被回收后，这些捕获的变量没有其他引用，会被GC回收。

	fmt.Println(nextInt()) // 1，i 不会被回收，而是逃逸到堆上，每次调用均可打印出来，而且访问的是同一个堆上的i
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInts := intSeq()
	fmt.Println(newInts()) // 1
	fmt.Println(newInts()) // 2

	fmt.Println(1<<10, 1024>>3, 1>>3) // 1024, 128, 0
}