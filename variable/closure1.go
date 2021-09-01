package main

import (
	"fmt"
)

func main1(){
	i := 1
    // add 1 in func
	func(){
		// 闭包获取变量，内部修改可以传到函数外
		i += 1
		fmt.Printf("1 in: %v\n", i)
	}()
	// 2
	fmt.Printf("2 out: %v\n", i)

	i = 1
	func(j int){
		// 形参修改，只在内部，不会传出函数，如果传入的是地址，也可以传到函数外
		j += 1
		fmt.Printf("3 in: %v\n", i)
	}(i)
	fmt.Printf("4 out: %v\n", i)

	i = 1
	func(j *int){
		// 形参修改，只在内部，不会传出函数，如果传入的是地址，也可以传到函数外
		*j += 1
		fmt.Printf("5 in: %v\n", i)
	}(&i)
	fmt.Printf("6 out: %v\n", i)

	// add in goroutine
	i = 1
	ch := make(chan int)
	go func() {
		i += 1
		ch <- 0
		fmt.Printf("7 in: %v\n", i)
		close(ch)

	}()
	<- ch
	// goroutine里的修改依然可以传出，此时没有值拷贝的限制，就可以理解成全局变量了
	fmt.Printf("8 out: %v\n", i)

}

func main(){
	main1()
	// fmt.Println(i)
	// 闭包变量作用域不会在这里的了
}