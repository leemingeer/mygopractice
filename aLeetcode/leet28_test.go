package main

import (
	"testing"
)


func TestStrStr(t *testing.T){

    tests := []struct{
		haystack string
		needle string
		expect int
	}{
    	{"hello", "ll", 2},
    	{"aaaaa", "bba", -1},
    	{"a", "a", 0},
    	{"", "", 0},
	}

	for _,tc := range tests{
		res := strStr(tc.haystack, tc.needle)
		if res != tc.expect{
			t.Errorf("got %d for input %s expects %d", res, tc.haystack, tc.expect)

		}

	}


}