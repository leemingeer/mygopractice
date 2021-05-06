package main

import "fmt"

func main(){
	var m = map[string]interface{}{"name": map[string]string{"sd": "li"},
		"location": "sh"}
	fmt.Println(len(m)) // 2

}
