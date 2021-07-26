package main

import "fmt"

func main() {
	slice1 := make([]int, 2, 10)
	fmt.Println(slice1, cap(slice1), len(slice1))

	slice1 = append(slice1, 1, 2)
	fmt.Println(slice1, len(slice1), cap(slice1))

	slice1 = append(slice1, []int{1, 2, 3, 4, 5, 6}...)
	fmt.Println(slice1, len(slice1), cap(slice1))

	slice1 = append(slice1, 7)
	fmt.Println(slice1, len(slice1), cap(slice1))

	slice2 := make([]int, 0, 10)
	fmt.Println(slice2, len(slice2), cap(slice2))

}
