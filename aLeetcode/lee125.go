package main

import (
	"fmt"
	"strings"
)
func isPalindrome125(ss string) bool {

	i,j:=0,len(ss)-1
	fmt.Println(i, j)
	for i<j{
		for{
			if skip(ss[i]){
				fmt.Println("skip")
				i++
				fmt.Println("i: ", i, string(ss[i]))
			} else{
				fmt.Println("break")
				break
			}
		}

		for{
			if skip(ss[j]){
				j--
				fmt.Println("j: ", j, string(ss[j]))
			} else{
				break
			}
		}

		fmt.Println("d:", string(ss[i]), string(ss[j]))

		if strings.ToLower(string(ss[i])) != strings.ToLower(string(ss[j])) {
			return false
		}
	}
	return true
}

func skip(c byte)bool{
	fmt.Println("in skip func: ", c, 'a'<=c)

	if ('a' <= c  && c <= 'z') || ('A' <= c && c <= 'Z'){
		return false
	}

	return true
}

