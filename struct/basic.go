package main

import (
	"fmt"
	"reflect"
)


type Person struct {
	age     int
	name string
}

var students = []Person{
	{ age: 2, name: "ming"},
	{ age: 2, name: "ming"},
	{ age: 3, name: "ming"},
	{ age: 1, name: "li"},
	{ age: 4, name: "li"},
	{ age: 3, name: "li"},
}



func main() {

	m2 := make(map[string]int)

    for _, p := range students{
    		m2[p.name]+= p.age
	}
	fmt.Println(m2)
    fmt.Println(reflect.TypeOf(students))



}
