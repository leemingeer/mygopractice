package main

import "fmt"

func main() {
	s := make([]int, 10)
	fmt.Printf("Addr of first element: %p\n", &s[0])
	fmt.Printf("Addr of slice itself:  %p\n", &s)
}
// why they are different?
// Addr of first element: 0xc00001e0a0
// Addr of slice itself:  0xc00000c060
