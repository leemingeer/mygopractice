package main

import "fmt"

func main() {

	var buffer [256]byte // byte零值是0, 256
	fmt.Println(10, byte(10))
	slice := buffer[10:20]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("before", slice)
	AddOneToEachElement(slice)
	fmt.Println("after", slice)
	fmt.Println("origin buffer", buffer)

	fmt.Println("Before: len(slice) =", len(slice))
	newSlice := SubtractOneFromLength(slice)
	fmt.Println("After:  len(slice) =", len(slice))
	fmt.Println("After:  len(newSlice) =", len(newSlice))
}

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}