package main

import (
	"math/rand"
)

func qsort(a []int, left,right int){
	if left > right{
		return
	}
	i := left
	j := right
	p := a[left]
	for {
		for a[j] > p{
			j --
		}
		for a[i] < p{
			i ++
		}

		if i == j {
			break
		}
		// 使用第一个元素作为a[pivot], 因为这句交换， a[i]的值才会动起来
		a[i], a[j] = a[j], a[i]
	}

	qsort(a, left, i-1)
	qsort(a, i+1, right)
}

func qsort2(a []int) []int {
	if len(a) < 2 { return a }
	left, right := 0, len(a) - 1
	// Pick a pivot
	pivotIndex := rand.Int() % len(a)
	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	//
	// left左变的一定比a[pivot]小，右边的一定比a[pivot]大
	// 最后再把a[pivot]跟a[left]交换
	// a[pivot]找到正确的位置
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	qsort2(a[:left])
	qsort2(a[left + 1:])

	return a
}