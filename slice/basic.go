package main

import "fmt"

func main() {

	var s1 []int
	fmt.Printf("s1: %v, %p\n", s1, &s1) //  [], 0xc00000c060
	fmt.Println(s1 == nil)              // true, nil表示没有引用内存
	fmt.Println(len(s1) == 0)           //true

	s2 := make([]int, 0)
	fmt.Printf("s2: %v, %p\n", s2, &s2) // [], 0xc00000c080
	fmt.Println(s2 == nil)              // false
	fmt.Println(len(s2) == 0)           //true

	s1 = append(s1, 1)
	fmt.Printf("%v\n", s1) //[1]
	
	s3:= []int{1,2}
	fmt.Println(s3, s3[2])
}
