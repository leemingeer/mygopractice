package main

import "fmt"

type T struct{}

func(t T) f(n int) T {
	fmt.Print(n)
	return t

}

func fx(n int){
	defer fmt.Print(n)
	n += 100
}

func main(){
	var t T
	defer t.f(1).f(2)
	fmt.Print(3)
	fx(4)
}
// 1342