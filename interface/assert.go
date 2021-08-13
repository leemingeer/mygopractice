package main

import "fmt"
func copy() int{}

func main(){
	var x interface{} = "foo"

	s  := x.(string)
	fmt.Println(s)     // "foo"

	n, ok := x.(int)
	fmt.Println(n, ok) // "0 false"

	nx := x.(int)        // ILLEGAL
	fmt.Println(nx)
	panic("x")
}