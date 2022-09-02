package main

import (
	"fmt"
)
func  main() {
	s := []int{4,5,6}
	fmt.Printf("%p\n", &s[0])
	s1 := s
	fmt.Print(s1)
	fmt.Printf("%p\n", &s1[0])
	// modify
	s[0] = 10
	fmt.Print(s1)

	s = nil
	fmt.Print(s1)
	s = append(s, 100)
	fmt.Print(s, s1)
}
