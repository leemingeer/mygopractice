package main

import (
	"math/rand"
)

func findKthLargest(nums []int, k int) int {

	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)

}

func quickSelect(nums []int, l, r, index int)int{
	p := randomPartition(nums, l, r)
	if p == index {
		return nums[p]
	} else if p < index {
		return quickSelect(nums, p+1, r, index)
	}
	return quickSelect(nums, l, p-1, index)

}


func randomPartition(nums []int, l, r int)int{
	pivot := rand.Int() % (r - l + 1) + l
	nums[pivot], nums[r] = nums[r], nums[pivot]
	x := nums[r]
	left := l

	for j :=l; j<r;j++{
		if nums[j] <= x{
			nums[j], nums[left] = nums[left], nums[j]
			left++

		}
	}
	nums[left], nums[r] = nums[r], nums[left]
	return left
}