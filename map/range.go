package main

import (
	"fmt"
	"sync"
)


type student struct {

	Age int
}

func generateData(data map[int]student){
	a := make([]int, size, size)
	for i := range a{
		data[i] = student{i} // data[i].Age = 1

	}
}

var loop int
var wg sync.WaitGroup
var size = 200
func main() {
	data := make(map[int]student, size)
	generateData(data)
	wg.Add(size)
	for i, s := range data{
		data[i] =  student{Age: 100 + s.Age}
		fmt.Println(i)
		go func(i int){
			defer wg.Done()
			loop += 1
			fmt.Println("in goroutine:", i)
		}(i)
	}
	wg.Wait()
	fmt.Println(loop)
}
