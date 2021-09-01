package main

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T){

	ts := []int{3,2,1,5,6,4}
	k := 2
	res := findKthLargest(ts, k)
	fmt.Println("res:", res)
	t.Logf("res: %d\n", res)
	if res == 5{
		t.Errorf("get res: %d\n", res)
	}

}
