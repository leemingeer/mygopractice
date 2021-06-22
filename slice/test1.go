package main

import "fmt"

func main() {
	a1 := make([]int, 10, 15)
	fmt.Println(&a1, a1)
	a2 := make([]int, 10, 15)
	fmt.Println(&a2, a2)
	//fmt.Println(a1==a2)
	//a := []int{7, 8, 9, 10}
	//b := a[15:16]
	//fmt.Println(a1)


	var s []int
	fmt.Println(s, s==nil)

	s = nil
	fmt.Println(s, &s, s==nil)

	s = []int(nil)
	fmt.Println(s, &s, s==nil)

	s = []int{}
	fmt.Println(s, &s, s==nil)

	s = make([]int, 0)
	fmt.Println(s, &s, s==nil)

	a3 := make([]int, 3)
	fmt.Println(len(a3), cap(a3), a3 ==nil, a3[0:], a3[1:], a3[2:], a3[3:])
}
