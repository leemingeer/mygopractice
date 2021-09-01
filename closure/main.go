package main

import "fmt"

func main(){
	var str = "abc"
	var s = []int{1,2,3}
	var m = map[string]int{"a": 1, "b":2}
	var ch = make(chan int)

	go func(){
		str = "def"
		s[0] = 100
		m["a"] = 1000
		ch <- 1
	}()

	// 修改都会传出
    fmt.Println(str)
	fmt.Println(s)
	fmt.Println(m)
	fmt.Println(<-ch)

}
