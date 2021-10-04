package main
import (
	"testing"
)

func TestIsPalindrome(t *testing.T){
	tcs := []struct{
		input string
		expect bool
	}{
		{"A man, a plan, a canal: Panama", true},
	}
	for _, tc := range tcs{
		res := isPalindrome125(tc.input)
		if res != tc.expect{
			t.Errorf("input %s, get %v, expect %v", tc.input, tc.expect, res)
		}
	}
}
