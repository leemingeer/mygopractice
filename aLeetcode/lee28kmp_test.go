package main

import "testing"

func TestStrstr(t *testing.T){
	tcs := []struct{
		haystack string
		needle string
		expect int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "", 0},
	}

	for _, tc := range tcs{
		res := strStrkmp(tc.haystack, tc.needle)
		if res != tc.expect{
			t.Errorf("haystack: %s, needle: %s, expect %d,  actual get: %d", tc.haystack, tc.needle, tc.expect, res)
		}
	}
}
