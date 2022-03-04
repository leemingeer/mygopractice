package main

import "fmt"

func heapify(heap *[]int, i int) {
	smallest := i
	lChild := 2*i + 1
	rChild := 2*i + 2

	if lChild < len(*heap) && (*heap)[lChild] > (*heap)[smallest] {
		smallest = lChild
	}
	if rChild < len(*heap) && (*heap)[rChild] > (*heap)[smallest] {
		smallest = rChild
	}

	if smallest != i {
		(*heap)[i], (*heap)[smallest] = (*heap)[smallest], (*heap)[i]
		heapify(heap, smallest)
	}
}

func buildHeap(heap []int) {
	n := len(heap)
	startIdx := n / 2 - 1

	for i := startIdx; i >= 0; i-- {
		heapify(&heap, i)
	}
}

func main(){
    h := []int{3,2,6,4,8,1,7,0}
	buildHeap(h)
    fmt.Println(h)
}