package main

import (
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
		{[]int{3,5,4,6,1,2}, "123456"},
		//{[]int{4,2,6,1,5,9,3}, "1234569"},
	}

	for _,tc := range tests{
		origin := make([]int, len(tc.input))
		origin2 := make([]int, len(tc.input))
		copy(origin, tc.input)
		copy(origin2, tc.input)
		qsort(tc.input, 0, len(tc.input)-1)
		if slice2string(tc.input) != tc.expect{
			t.Errorf("got %v for input %v expects %s", tc.input, origin, tc.expect)
		}
		qsort2(origin)
		if slice2string(tc.input) != tc.expect{
			t.Errorf("got %v for input %v expects %s", tc.input, origin2, tc.expect)
		}

	}


}