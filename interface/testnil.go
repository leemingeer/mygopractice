package main

import (
	"fmt"
	"reflect"
)

func main(){

	//var a = (*int)(unsafe.Pointer(uintptr(0x0)))
	//print(a == nil)

	type A interface{}
	type B struct{}
	var a *B

	bb := (*B)(nil)
	fmt.Println(reflect.TypeOf(bb))

	print(a == nil)            //true
	print(a == (*B)(nil))      //true
	fmt.Println("x:", (*B)(nil) == nil) // true
	print((A)(a) == (*B)(nil)) //true

	print((A)(a) == nil)       //false
	fmt.Println(reflect.TypeOf((A)(a)))
}