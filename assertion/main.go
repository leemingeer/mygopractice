package main

import "fmt"

func main(){
	var x interface{} = "ming"
	v := x.(string)
	fmt.Println(v)

	var x1 interface{}
	v1:= x1.(string)
	fmt.Println(v1)
}
