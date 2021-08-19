package main

import "fmt"

// can
func appendSiceP(s *[]int){
	*s = append(*s, 4)
}

// can't
func appendSice(s []int){
	s = append(s, 4)
}
// 只有值传递
// append修改了slice的length属性，会返回一个新类型的slice

func modifySiceLen(s []int){
	// modify slice start point and length, return a new type slice
	s = s[1:2]
}
// can
func modifyElementSice(s []int){
	s[0] = 10
}



func main(){
	arr := []int{1, 2 ,3}
	appendSice(arr)
	fmt.Println(arr)

	modifySiceLen(arr)
	fmt.Println(arr)

	appendSiceP(&arr)
	fmt.Println(arr)

	modifyElementSice(arr)
	fmt.Println(arr)
}
