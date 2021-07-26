package main

import "fmt"

type Person struct {
	age     int
	address []string
}

func main() {

	var one Person
	fmt.Printf("one: %v,%v, %p\n", one, &one, &one) // {0 []},&{0 [], 0xc0000a6020

	two := Person{}
	fmt.Printf("two: %v,%v, %p\n", two, &two, &two) //{0 []},&{0 []}, 0xc0000a6080

	three := Person{age: 12,
		address: []string{"shanghai", "SD"},
	}

	fmt.Printf("three: %v,%v, %p\n", three, &three, &three) //{12 [shanghai SD]},&{12 [shanghai SD]}, 0xc00000c0c0

}
