package main

import "fmt"

type city struct{
	location string
}

func main(){

	a1 := city{}
	fmt.Println(a1.location)

	var a2 city
	fmt.Println(a2.location)

}