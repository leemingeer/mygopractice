package main

import "fmt"

var cache = map[int]int{}

func lengthOfLIS(nums []int) int{
	maxLength := 1
	for i:=0;i<len(nums);i++{
		// 这里返回当前i索引位置的最长递增子序列
		maxLength = max(maxLength, L(nums, i))
	}
	return maxLength
}

func max(a,b int)int{
	if b > a {
		return b
	}
	return a
}

func L(nums []int, i int)int{

	if v, ok := cache[i]; ok{
		return v
	}
	// 最后一个元素没有可以和它自身组成自增序列了，直接返回1表示自身
	// 这里要主动返回，如果不主动返回，靠for循环结束完，则返回的是返回值int类型的0值
    //if i == len(nums) - 1{
    //	return 1
	//}
	tmpMaxLength := 1
	for j:=i+1;j<len(nums);j++{
		if nums[j] > nums[i]{
			// 这里的1是指，i和j两个元素构成的子序列 + 以j开始的元素的最长递增子序列
			tmpMaxLength = max(tmpMaxLength, L(nums, j)+1)
		}
	}
	// 引入缓存，这里有些类似dp了
	cache[i] = tmpMaxLength
	return tmpMaxLength
}

func main(){

	// 8
	// nums := []int{10,9,2,5,3,7,101,18}

	//nums := []int{0,1,0,3,2,3}
    nums := []int{7,7,7,7,7,7,7}
	fmt.Println(lengthOfLIS(nums))

}