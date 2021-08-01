package main

import (
	"fmt"
	"reflect"
)

func main() {
	fVar := 3.14
	v := reflect.ValueOf(fVar)
	// 指针类型才可set
	fmt.Printf("is float canSet: %v canAddr %v\n", v.CanSet(), v.CanAddr())
	//v.SetFloat(221232.12)
	vp := reflect.ValueOf(&fVar) // Value指向一个point
	// elem(), 对这个point解引用
	fmt.Printf("is float canSet: %v canAddr %v\n", vp.Elem().CanSet(), vp.Elem().CanAddr())
	vp.Elem().SetFloat(2.33333)

	println(fVar)
}
