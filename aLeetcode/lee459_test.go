package main

import "testing"

func TestRepeatstring(t *testing.T){
	tcs := []struct{
		input string
		expect bool
	}{
		{"abab", true},
		{"aba", false},
		{"abcabcabcabc", true},
	}

	for _, tc := range tcs{
		res := repeatedSubstringPattern(tc.input)
		if res != tc.expect{
			t.Errorf("input: %s, expect %v,  actual get: %v", tc.input, tc.expect,  res)
		}
	}
}