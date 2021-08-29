package main

import (
	"fmt"
	"math/rand"
)

func qsort(a []int, left,right int){
	if left > right{
		return
	}
	i := left
	j := right
	p := a[left]
	for i < j{
		for a[j] > p{
			j --
		}
		for a[i] < p{
			i ++
		}
		a[i], a[j] = a[j], a[i]
	}

	qsort(a, left, i-1)
	qsort(a, i+1, right)
}

func qsort2(a []int) []int {
	fmt.Println("---next level---", a)
	if len(a) < 2 { return a }

	left, right := 0, len(a) - 1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)
	fmt.Println("pivotIndex: ", pivotIndex)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		fmt.Println(a, i, left, right)
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	fmt.Println("after for:", left, right)

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	qsort2(a[:left])
	qsort2(a[left + 1:])

	return a
}