package main

import "fmt"

func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}
	next := make([]int, len(s))
	getNext459(s, next)
	if next[len(s)-1] != -1 && len(s) %(len(s) -(next[len(s) - 1] + 1)) == 0{
		return true
	}

	return false
}

func getNext459(s string, next []int){
	fmt.Println(s)
	j := -1
	next[0] = j
	for i:=1;i<len(s);i++{
		for ;j>=0 && s[i] != s[j+1];{
			j = next[j]
		}
		if s[i] == s[j+1]{
			j++
		}
		next[i] = j
	}
	fmt.Println(next)
}