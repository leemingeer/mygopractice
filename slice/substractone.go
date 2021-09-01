package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("%p\n", &a)
	a = 20
	fmt.Printf("%p\n", &a)


	var buffer [256]byte // byte零值是0, 256
	slice := buffer[10:20]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("Before: len(slice) =", len(slice))
	newSlice := SubtractOneFromLength(slice)
	fmt.Println("After:  len(slice) =", len(slice))
	fmt.Println("After:  len(newSlice) =", len(newSlice))
	newSlice[0] = 100
	fmt.Println(newSlice)

	slice3 := append([]byte(nil), newSlice...)
	fmt.Println("Copy a slice:", slice3)
}

func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}