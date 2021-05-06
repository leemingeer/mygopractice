package main

import (
	"fmt"
	"reflect"
)

func main(){
	var s1 =[]string{}
	for s := range s1{
		fmt.Println(s)
	}

	var s2 =[]interface{}{}
	for s := range s2{
		fmt.Println(reflect.TypeOf(s))
	}
}