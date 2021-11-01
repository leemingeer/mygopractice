package main

import "fmt"

var resint int
func lengthOfLIS(nums []int) int {
	path := []int{nums[0]}

	bt(nums, 0, path)
	return resint

}

func bt(nums []int, start int, path []int){
	fmt.Println("next: ", start, len(nums) - 1)
	if start >= len(nums) - 1{
		fmt.Println("2:", resint, path)
		if resint < len(path){
			fmt.Println(path)
			resint = len(path)
		}
		return
	}

	for i:= start;i<len(nums);i++{
		for j:=i;j<len(nums);j++{
			if nums[j] > nums[i]{
				path = append(path, nums[j])
				fmt.Println("1:",i, j, path)
				bt(nums, start+1, path)
				path = path[:len(path)-1]
			}
		}
	}
}

func main(){
	nums := []int{1,5,3}
	fmt.Println(lengthOfLIS(nums))
}