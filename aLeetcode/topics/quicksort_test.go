package main

import (
	"fmt"
	"strconv"
	"testing"
)

func slice2string(a []int)string{
	res := ""
	for _, v := range a {
		res += strconv.Itoa(v)
	}
    return res

}

func TestQsort(t *testing.T){

	tests := []struct{
		input []int
		expect string
	}{
		{[]int{6,5,4,3,2,1}, "123456"},
		//{[]int{4,2,6,1,5,9}, "124569"},
	}

	for _,tc := range tests{
		origin := make([]int, len(tc.input))
		copy(origin, tc.input)
		//qsort(tc.input, 0, len(tc.input)-1)
		//if slice2string(tc.input) != tc.expect{
		//	t.Errorf("got %v for input %v expects %s", tc.input, origin, tc.expect)
		//}
		ret := qsort2(tc.input)
		fmt.Println("result: ", ret)
		if slice2string(tc.input) == tc.expect{
			t.Errorf("got %v for input %v expects %s", tc.input, origin, tc.expect)
		}

	}


}