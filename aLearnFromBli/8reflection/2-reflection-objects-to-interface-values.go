package main

import (
	"fmt"
	"reflect"
)

// 可逆的

type student struct {
	Name string
}

func main() {
	// FIRST EXAMPLE SHOWING CONVERT REFLECT.VALUE TO FLOAT
	floatVar := 3.14
	v := reflect.ValueOf(floatVar)
	// v.Interface() 返回潜在类型是v的interface{}，var i interface{} = (v's underlying value)
	// 重新构建float
	newFloat := v.Interface().(float64)
	fmt.Println(newFloat + 1.2)

	// second example showing convert Reflect.Value to slice
	sliceVar := make([]int, 5)
	v = reflect.ValueOf(sliceVar)
	// 方式1. 在Value形态下，reflect.append入参就是Value类型的
	v = reflect.Append(v, reflect.ValueOf(2))
	newSlice := v.Interface().([]int)
	// 方式2，在slice形态下，append
	newSlice = append(newSlice, 4)
	fmt.Println(newSlice)

	// third example showing convert Reflect.Value to student
	// reflect.New返回指向zero Value的指针， 跟ming不ming没关系
	stuPtr := reflect.New(reflect.TypeOf(student{Name: "ming"}))
	stu := stuPtr.Elem()
	fmt.Println(stu, reflect.TypeOf(stu))
	nameField := stu.FieldByName("Name")
	if nameField.IsValid() {
		if nameField.CanSet() {
			nameField.SetString("chong")
		}
		realStudent := stu.Interface().(student)
		fmt.Println(realStudent)
	}

}
