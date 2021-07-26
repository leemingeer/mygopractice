package main

import "fmt"

func main() {
	d2slice := make([][]int, 2, 10)
	fmt.Println(d2slice, len(d2slice), cap(d2slice)) //[[] []] 2 10
	d2slice = append(d2slice, []int{1, 2})
	d2slice = append(d2slice, []int{3, 4})
	fmt.Println(d2slice, len(d2slice), cap(d2slice)) //[[] [] [1 2] [3 4]] 4 10
}
