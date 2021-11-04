package main

import "fmt"

func main(){
	var s1 = "1T"
	var s2 = "2T"

	if s1 < s2{
		fmt.Printf("%s is less than %s", s1, s2)
	} else{
		fmt.Printf("%s is bigger than %s", s1, s2)
	}
}
