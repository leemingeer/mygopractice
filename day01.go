package main

import (
	"fmt"
)

func main() {

	defer_call()

}

func defer_call() {
	defer func() { fmt.Println("before") }()
	defer func() { fmt.Println("middle") }()
	defer func() { fmt.Println("after") }()
	panic("exception")
}
