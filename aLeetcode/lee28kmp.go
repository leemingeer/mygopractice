package main

func strStrkmp(haystack string, needle string) int {

	next := make([]int, len(needle))

	if len(needle) == 0 {
		return 0
	}
	j := 0
	getNext(next, needle)
	for i := 0; i < len(haystack); i++ {
		for ; j >= 0 && haystack[i] != needle[j]; {
			j = next[j]
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j == len(needle)-1 {
			return i - len(needle) + 1
		}
	}
	return -1
}

func getNext(next []int, needle string){
	j := -1
	next[0] = -1
	for i:= 1;i<len(needle);i++{
		for;j>=0 && needle[i] != needle[j+1];{
			j = next[j]
		}
		if needle[i] == needle[j+1]{
			j++
		}
		next[i] = j
	}
}