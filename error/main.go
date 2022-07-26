package main

import (
	"fmt"
	"reflect"
)

type Person struct{
	name string
}

func (p *Person)SayHi(){
	fmt.Println("Hi")
}

func (p *Person) Name() string{
	if p == nil{
		return "nil receiver"
	}
	return p.name
}


func main() {
	var p *Person
	p.SayHi() // Hi
	fmt.Println(p.Name())

	fmt.Println(reflect.TypeOf(nil), nil)
}
