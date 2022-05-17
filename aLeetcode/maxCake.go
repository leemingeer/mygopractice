package main

import "fmt"

func maxCakePrice(n int, cakeprice []int, cache map[int]int)int{

	if n == 1{
		return cakeprice[n]
	}

	if n < 1{
		return 0
	}


	if _, ok := cache[n]; ok {
		return cache[n]
	}
	tmpMax :=0
	for i:=1;i<7;i++{

		tmpMax = max1(tmpMax,  maxCakePrice(n-i, cakeprice, cache) + cakeprice[i])
	}
	cache[n] = tmpMax
	return tmpMax

}


func max1(a, b int)int{
	if b > a{return b}
	return a
}

func main(){
	cakeprice := []int{0, 2, 3, 6, 7, 11, 15}
    cache := make(map[int]int, 8)

	fmt.Println(maxCakePrice(30, cakeprice, cache))

}