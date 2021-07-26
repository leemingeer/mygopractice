package main

import "fmt"

func main() {
	const (
		a, b = "golang", 100
		d, e
		f bool = true
		g
	)
	fmt.Print(d, e, g)

}

// golang100 true
