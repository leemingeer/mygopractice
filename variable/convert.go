package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main(){
	// Convert string to bytes
	b := []byte("ABC€")
	fmt.Println(b) // [65 66 67 226 130 172]
	// Convert bytes to string
	s := string([]byte{65, 66, 67, 226, 130, 172})
	fmt.Println(s) // ABC€


	sp := strings.Split("a,b,c", ",")
	fmt.Println(reflect.TypeOf(sp)) // []string
	fmt.Println(sp) // [a b c]
	fmt.Printf("%#v\n", sp) // []string{"a", "b", "c"}
    fmt.Printf("%q\n", sp) // ["a" "b" "c"]

}