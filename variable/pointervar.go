package main

import (
	"fmt"
)

var p *int

func foo() (*int, error) {
	var i = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Printf("3: %p, %T\n", p, p)
	fmt.Println(*p)
}

func main() {
	var err error
	fmt.Printf("1: %p, %T, %d\n", p, p, *p)
	p, err = foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("2: %p, %T\n", p, p)
	bar()
	fmt.Println(*p)
}
