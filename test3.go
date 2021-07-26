package main

import "fmt"

func f1() {
	var a int8 = -1
	var b int8 = -128 / a
	fmt.Println(b)
	// -128

}

//func f2(){
//	const c int8 = -1
//	var d int8 = -128/c
//	fmt.Println(d)
//	// constant 128 overflows int8
//
//}

func main() {
	f1()
	// f2()
}
