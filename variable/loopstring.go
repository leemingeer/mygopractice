package main

import "fmt"

var ss = "you are a main"

func main(){
	for _, v := range ss{
		fmt.Println(string(v))
	}
}
