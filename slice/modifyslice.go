package main

import "fmt"

func main() {
	a := make([]int, 2, 5)
	fmt.Println(a) // [0, 0]
	for _, v := range a{
		a[1] += 2
		v += 1
		fmt.Println(v)
	}
	fmt.Println(a)

	fmt.Println("--------")
	m := make(map[int]int, 2) // map[]
	fmt.Println(m)
	for k, v := range m{
		m[k] += 2
		fmt.Println(v)
	}
	fmt.Println(a)

}