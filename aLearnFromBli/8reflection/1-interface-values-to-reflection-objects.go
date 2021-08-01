package main

import (
	"fmt"
	"reflect"
)

func printMeta(obj interface{}) {
	// 接口变量都有 pair: <value, type>
	t := reflect.TypeOf(obj) // 返回type接口类型变量
	n := t.Name()
	k := t.Kind()
	v := reflect.ValueOf(obj)
	fmt.Printf("Type: %s Type.Name: %s Kind: %s Value: %v\n", t, n, k, v)
}

type handler func(int, int) int // 输入是两个int， 返回是一个int，这种类型的func，函数实现不确定，后面可以具体赋值

func main() {

	var intVar int64 = 10
	stringVar := "hello"
	type book struct {
		name  string
		pages int
	}
	sum := func(a, b int) int {
		return a + b
	}
	var sub handler = func(a, b int) int {
		return a - b
	}
	sli := make([]int, 5)
    // 原始类型， Type, TYpe.Name, Type.Kind都是一样的
	printMeta(intVar)
	printMeta(stringVar)


    // Type: main.book, main包中我们定义的类型
    // Type.Name, 在包中我们定义的名字，即book，如果没有type xxx 去重新定义， 则为空""
    // Type.Kind，最本质的类型，即编译器直到的类型
	printMeta(book{
		name:  "harry potter",
		pages: 500,
	})


	printMeta(sum)
	printMeta(sub)
	printMeta(sli)
}
