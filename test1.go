package main

import "fmt"

type T struct{}

func (t T) f(n int) T {
	fmt.Println(n)
	return t

}

func fx(n int) {
	defer fmt.Println(n)
	n += 100
	fmt.Println(n)
}

func ret() {
	fmt.Println("ret function")
	//return
	fmt.Println("after return")
}

func main() {
	var t T
	defer t.f(1).f(2)
	fmt.Println(3)
	fx(4)

	ret()
}
