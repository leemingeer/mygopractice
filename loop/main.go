package main

import (
	"fmt"
)

var global int

func main(){
	arrays := []int{1,2,3,4}

	var finalarrays []int

	for _, v := range(arrays){
		finalarrays = append(finalarrays, v)
	}
	fmt.Println(finalarrays)

	type A struct {
		a int
	}

	b := []A{
		{a: 2},
		{a: 4},
		{a: 6},
	}

	bb := make([]A, 0) // 这里slice里的元素类型是指针，如果是值就没有问题了
	for _, a := range b {
		//a := a // 取地址，
		//
		bb = append(bb, a)
	}
	fmt.Println(bb)


	//var finalarrayspoint []*int
	//for _, v := range(arrays){
	//	finalarrayspoint = append(finalarrayspoint, &v)
	//}
	//for _, v := range(finalarrayspoint){
	//	fmt.Println(*v)
	//}
	//
    //// 容易出错的闭包问题
    //// i 作为参数传入
	//for i:=0; i<10; i++{
	//
	//	go func(){
	//		fmt.Println(i)
	//	}()
	//}
	//time.Sleep(time.Second)
	//
	//for i:=0; i<10; i++{
	//
	//	go func(i int){
	//		global := doublei(i)
	//		fmt.Println(i, global)
	//
	//	}(i)
	//}
	//time.Sleep(time.Second)

}

func doublei(i int) int{
	global = 2 * i
	return global
}