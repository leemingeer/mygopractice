package main

import (
	"testing"
)

func TestLong(t *testing.T) {
	tcs := []struct{
		Input string
		expect struct{
			num int
			s string
		}
	}{
		{"aaa", struct{num int; s string}{3, "aaa"}},
		{"aaaa",  struct{num int; s string}{4, "aaaa"}},
		{"xyaabaaaabbaaz",  struct{num int; s string}{6, "baaaab"}},

	}

	for _, tc := range tcs {
		length, res := longestPara(tc.Input)
		if length != tc.expect.num || res != tc.expect.s{
			t.Errorf("get length:%d string :%s expect %d,  %s", length, res, tc.expect.num, tc.expect.s)
		}
	}


}