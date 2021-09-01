package main

import "fmt"

func combine(n int, k int) [][]int {
	res := [][]int{}

	backtrack(1, k, n, []int{}, &res)

	return res

}

func backtrack(start, k, n int, track []int, res *[][]int) {
	fmt.Println("---------")

	if len(track) == k {
		//temp := make([]int, k)
		//copy(temp, track)
		*res = append(*res, track)
		fmt.Printf("res: %v, addr: %v, len: %d, cap: %d\n", *res, &(*res)[0][0], len(*res), cap(*res))
		fmt.Printf("track: %v, addr: %v, len: %d, cap: %d\n", track, &track[0], len(track), cap(track))

		return
	}
	for i := start; i <= n; i++ {
		fmt.Printf("i: %d\n", i)
		track = append(track, i)
        fmt.Printf("track1: %v\n", track)
		backtrack(i+1, k, n, track, res)
		fmt.Printf("track2: %v\n", track)
		track = track[:len(track)-1]
	}
}

func main(){
	fmt.Println(combine(4, 2))
}