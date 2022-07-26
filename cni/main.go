package main

import (
	"fmt"
	"os"
)

type T int32

func (t T) Error() string {
	return "bad error"
}

func printEmptyInterfaceAndNonEmptyInterface() {
	var eif interface{} = T(5)
	var err error = T(5)
	println("eif:", eif)
	println("err:", err)
	println("eif = err:", eif == err)

	err = T(6)
	println("eif:", eif)
	println("err:", err)
	println("eif = err:", eif == err)
}

type Worker interface{
	Work() error
}

type Qstruct struct{}

func (q Qstruct) Work() error{
	return nil
}

func findSomething() Worker{
	return nil
}

func main(){
	var v Worker
	*v = findSomething()
	if v != nil{
		fmt.Printf("v is not nil: %#v\n", v)
		println(v)
		os.Exit(0)
	}
	println(v)
}