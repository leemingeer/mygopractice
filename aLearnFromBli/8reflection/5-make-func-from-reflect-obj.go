package main

import (
	"fmt"
	"reflect"
	"time"
)

func makeTimeFunc(f interface{}) interface{} {
	tf := reflect.TypeOf(f)
	vf := reflect.ValueOf(f)

	if tf.Kind() != reflect.Func {
		panic("expect a function")
	}
	// par1: reflect.Type类型
	// par2: 看定义是传入函数类型，比如输入两个int，返回一个float的函数类型，实际调用是函数实现。
	wrapper := reflect.MakeFunc(tf, func(args []reflect.Value) (result []reflect.Value) {
		start := time.Now()
		result = vf.Call(args)
		end := time.Now()
		fmt.Printf("The function takes %v\n", end.Sub(start))
		return result
	})
	return wrapper.Interface() // Value 变成Interface
}

// 这个函数的类型就是func(), 没有入参，没有返回值
func TimeMe() {
	time.Sleep(1 * time.Second)
}

func main() {
	timedFunc := makeTimeFunc(TimeMe).(func())// interface 专成 func()）类型
	timedFunc()
}