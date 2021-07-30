package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string
	Age int
}



func main(){
	a := map[string]struct{}{"x": {}}
	fmt.Println(a, reflect.TypeOf(a))


	var zs student
	var ls = student{}
	var xiaoming = student{"ming", 22}

	fmt.Printf("%+v\n", zs)
	fmt.Printf("%+v\n", ls)

	var class1 []student

	class1 = append(class1, xiaoming)
	for _, v := range class1{
		fmt.Printf("range class1 %+v", v)
	}


}
