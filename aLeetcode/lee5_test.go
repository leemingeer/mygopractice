package main

import "testing"

func TestPa(t *testing.T){
	tcs := []struct {
		input  string
		expect string
	}{
		{"babad", "bab" },
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
	}

	for _, tc := range(tcs){
		res := longestPalindrome(tc.input)
		if res != tc.expect{
			t.Errorf("input: %s, expect: %s, get: %s", tc.input,tc.expect, res)
		}
	}



}
