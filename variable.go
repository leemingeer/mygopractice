package main

import (
	"fmt"
	"reflect"
)

var name string

func setName(){
	name = "Li"
	valueofname := reflect.ValueOf(name)
	fmt.Println(1, valueofname.Interface())
	fmt.Println(2, valueofname.String())

	valueOfName := reflect.ValueOf(&name)
	valueOfName.Elem().Set(reflect.ValueOf("Mi"))
	fmt.Println(3, name)


}

func showName(){
	fmt.Println(4, name)
}

func showvar(){
	var aa string // local variable
	aa ="varinfunc"
	fmt.Println("aa in shwovar:", aa)
}


func main(){
	fmt.Printf("name: %s\n", name)
	setName()
	showName()
	fmt.Printf("name: %s\n", name)
	showvar()

}