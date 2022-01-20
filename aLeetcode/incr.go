package main

import "fmt"

var path []int
func inc(nums []int, i int){
    fmt.Println(i)
	for j:=i;j<len(nums);j++{
		if nums[j] > nums[i]{
			path = append(path, nums[j])
			fmt.Println("path: ", path)
			inc(nums, j)

		}
	}
}


func main(){
	nums := []int{3,2,6,9,1,4,7}
	inc(nums, 0)
	fmt.Println(path)

}
