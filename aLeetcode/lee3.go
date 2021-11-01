package main

func lengthOfLongestSubstring(s string) int {
	res := ""
	for i,j := 0,0; j<len(s);j++{
		for k:=i; k<j; k++{
			if s[k] == s[j]{
				i = k+1
			}
		}
		if j - i + 1 > len(res){
			res = s[i:j+1]
		}
	}
	return len(res)
}