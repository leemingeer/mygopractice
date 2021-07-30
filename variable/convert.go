package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"github.com/inhies/go-bytesize"
)

func main() {
	// Convert string to bytes
	b := []byte("ABC€")
	fmt.Println(b) // [65 66 67 226 130 172]
	// Convert bytes to string
	s := string([]byte{65, 66, 67, 226, 130, 172})
	fmt.Println(s) // ABC€

	sp := strings.Split("a,b,c", ",")
	fmt.Println(reflect.TypeOf(sp)) // []string
	fmt.Println(sp)                 // [a b c]
	fmt.Printf("%#v\n", sp)         // []string{"a", "b", "c"}
	fmt.Printf("%q\n", sp)          // ["a" "b" "c"]

	b1, _ := bytesize.Parse("1024GB")
	b1_format := b1.Format("%.2f ", "byte", true)
	fmt.Printf("%s\n", b1_format)

	b2, _ := bytesize.Parse("1023 GB")
	fmt.Printf("%s\n", b2)

	b3, _ := bytesize.Parse(" 10 TB")
	b3_format := b3.Format("%.2f ", "byte", true)
	fmt.Printf("%s\n", b3)
	fmt.Printf("%s\n", b3_format)

	fmt.Println(b1_format < b3_format)
	b4, _ := strconv.ParseBool("False")
	fmt.Println(b4, b4 == false)
	iopsLimitInt, err := strconv.Atoi("")
	if err != nil {
	}
	fmt.Printf("result: %d", iopsLimitInt)


}