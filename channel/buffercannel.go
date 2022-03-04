package main

import (
"fmt"
)


func main() {
	ch := make(chan string, 0)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<- ch)
	fmt.Println(<- ch)
}