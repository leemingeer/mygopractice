package main

func strStr(haystack string, needle string) int {

	if len(needle) == 0{
		return 0
	}
	lengthhaystack := len(haystack)
	lengthneedle := len(needle)

	if lengthneedle > lengthhaystack{
		return -1
	}

	for i:=0; i<= lengthhaystack - lengthneedle; i++{
		if haystack[i: i+lengthneedle] == needle{
			return i
		}
	}
	return -1
}