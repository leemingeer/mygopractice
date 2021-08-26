package main

import "fmt"

var resTest [][]int

func main(){
	track := []int{1,2}
	backtrackTest(track)
	track = track[:len(track)-1]
	// append的是slice， slice是地址，没有发生扩容，地址是同一个，导致前面写入的元素发生改变，不是期望的，在所容时是好使用copy将期望拷贝到一块独立的内存，避免原地值的slice被修改。
	track = append(track, 10)
	// 发生扩容，第一次append的slice的地址是扩容前的，没有被修改。所以是我们期望的
	//track = append(track, 10, 11)
	backtrackTest(track)
}

func backtrackTest(track []int) {
	res = append(resTest, track)
	fmt.Println("res: ", res)
}


