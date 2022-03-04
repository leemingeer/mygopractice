package main

import (
	"fmt"
	"reflect"
)

func main(){
	var s1 =[]string{}
	for s := range s1{
		fmt.Println("in loop")
		fmt.Println(s)
	}
	fmt.Println("mark")

	var s2 =[]interface{}{}
	for s := range s2{
		fmt.Println(reflect.TypeOf(s))
	}

	var s3 []string
	s3 = append(s3, "ming")
	fmt.Println(s3)
}