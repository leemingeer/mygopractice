package main

import "fmt"

func addBaseV1(base int) func(i int) int {
	return func(i int) int {
		base += i
		return base
	}

}

func addBaseV2() func(i int) int {
	base := 100
	return func(i int) int {
		base += i
		return base
	}

}

func cal(base int)(func(int)int, func(int)int){
	add := func(i int)int{
		base += i
		return base
	}

	sub := func(i int) int{
		base -= i
		return base
	}

	return add, sub
}

func main(){
	f1, f2 := cal(10)
	fmt.Println(f1(1), f2(2)) // 11, 9
	fmt.Println(f1(3), f2(4)) // 12, 8
	fmt.Println(f1(5), f2(6)) // 13, 7
}

//func main() {
//	addV1 := addBaseV1(100)
//	fmt.Println("result: ", addV1(10))
//	fmt.Println("result: ", addV1(30))
//
//	addV2 := addBaseV2()
//	fmt.Println("result: ", addV2(10))
//	fmt.Println("result: ", addV2(30))
//}
