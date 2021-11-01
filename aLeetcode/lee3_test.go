package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T){
	tcs := []struct {
		input  string
		expect int
	}{
		{"abcabcbb", 3 },
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, tc := range(tcs){
		res := lengthOfLongestSubstring(tc.input)
		if res != tc.expect{
			t.Errorf("input: %s, expect: %d, get: %d", tc.input,tc.expect, res)
		}
	}



}