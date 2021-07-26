package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4}
	m := make(map[int]*int)
	for k, v := range s {
		n := v
		fmt.Printf("n addr:%v, value: %d\n", &n, n)

		m[k] = &n
	}
	for key, value := range m {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
	fmt.Println(m)

	for _, v := range []int{2, 2, 2, 2, 2} {
		x := v // x addr differs even when v is always 2
		fmt.Printf("x addr:%v, value: %d\n", &x, x)
	}
}
