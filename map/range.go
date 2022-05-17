package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4}
	m := make(map[int]int)
	for k, v := range s {
		//n := v
		m[k] = v
	}
	for key, value := range m {
		fmt.Printf("map[%v]=%v\n", key, value)
	}
	fmt.Println(m)

	for _, v := range []int{2, 2, 2, 2, 2} {
		x := v // x addr differs even when v is always 2
		fmt.Printf("x addr:%v, value: %d\n", &x, x)
	}

	tmpMap := make(map[string]int)
	// var tmpMap map[string]int
	fmt.Println(tmpMap["ming"])
	tmpMap["ming"] = 1
	fmt.Println(tmpMap["ming"])
}
