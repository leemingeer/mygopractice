package main

import (
	"fmt"
)

var x = 1

func change() (int,error){
	return 3, nil
}

func main() {
	var err error
	fmt.Printf("%p\n", &x)
	x, err = change()
	fmt.Printf("%p\n", &x)
	if err != nil{
		fmt.Println(err)
	}
	func(){
		fmt.Printf("show1: %d, %p\n", x, &x)
	}()

	show()
	fmt.Printf("outer: %d, %p\n", x, &x)

	var ifx = 1
	if true{
		ifx = 2
	}
	fmt.Println("ifx:", ifx)
	
	var forx = 1
	for i:=0;i<1;i++{
		fmt.Println("i", i)
		forx = 2
	}
	fmt.Println("forx:", forx)
}

func show(){
	fmt.Printf("show2: %d, %p\n", x, &x)
}
