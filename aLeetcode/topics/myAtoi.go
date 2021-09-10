package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main(){
	i := 41624
    var res []int
	// int to int slice
    for i > 0{
    	rest := i % 10
        res = append(res, rest)
    	i = i / 10
	}
	fmt.Println(res)

	// string to int
	istr := "041624"
	res2 := 0
	strconv.Atoi()
	fmt.Println("12", reflect.TypeOf("12")) // 12, string, 对应的是int32: 48
	for _, c := range istr{
		fmt.Println(int(c), reflect.TypeOf(c)) // 48, int32
		res2 = res2 * 10 +  int(c - '0')
	}
	fmt.Println(res2)



}
