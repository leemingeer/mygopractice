package main

import "fmt"

var res[][]int

func combine(n int, k int) [][]int {
	res = [][]int{}
	backtrack(1, k, n, []int{})
	return res
}

func backtrack(start, k, n int, track []int) {

	if len(track) == k {
		temp := make([]int, k) // 申请len，cap为k的底层数组，slice指向全组元素
		copy(temp, track) // right ro left
		// 这里必须先copy一份内存，因为slice是地址传递，append的时候是地址，而在回溯的时候，还是原先那个地址的底层数组，但是值会被修改，导致之前append进来的slice指向的底层数组被修改了
		res = append(res, temp)
		return
	}
	for i := start; i <= n; i++ {
		track = append(track, i)
		backtrack(i+1, k, n, track)
		track = track[:len(track)-1]
	}
}

func main(){
	fmt.Println(combine(4, 2))
}