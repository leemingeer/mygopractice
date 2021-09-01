package main

import (
	"fmt"
)
func  main() {
	s := []int{4,5,6}
	fmt.Printf("%p\n", &s[0])
	s1 := s[:1]
	s2 := s[:2]
	fmt.Printf("%p , %p\n", s1, s2)

	fmt.Printf("%p\n", &s1)

}
