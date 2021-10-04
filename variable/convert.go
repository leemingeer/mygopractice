package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	fmt.Println('A') // 65
	fmt.Println("A") // "A"
	fmt.Println(byte('A'), []byte("A")) // 65, [65]

	fmt.Println(6, '6', "6")
	fmt.Println("54", string(54))

	fmt.Println(string(65)) // "A"
	fmt.Println(byte(65)) // 65
	fmt.Println( "66") // 65

	// Convert string to bytes
	b := []byte("ABC")
	fmt.Println(b) // [65 66 67]
	// Convert bytes to string
	s := string([]byte{65, 66, 67})
	fmt.Println(s) // ABC



	sp := strings.Split("a,b,c", ",")
	fmt.Println(reflect.TypeOf(sp)) // []string
	fmt.Println(sp)                 // [a b c]
	fmt.Printf("%#v\n", sp)         // []string{"a", "b", "c"}
	fmt.Printf("%q\n", sp)          // ["a" "b" "c"]

	b4, _ := strconv.ParseBool("False")
	fmt.Println(b4, b4 == false)
	iopsLimitInt, err := strconv.Atoi("")
	if err != nil {
	}
	fmt.Printf("result: %d", iopsLimitInt)


}