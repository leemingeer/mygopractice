package main

import (
	"fmt"
)

func main() {
	var x = 1
	if true {
		x = 2
		y := 3 // y作用域仅在if内部
		fmt.Println(y)
	}
	fmt.Printf("inner x: %d\n", x)
	//fmt.Printf("inner y: %d\n", y)
	fmt.Println(x)
}
