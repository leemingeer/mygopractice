package main

import "fmt"

type IntSet struct{}


func(s IntSet) string(){
	fmt.Println("OK！")
}

func(s *IntSet) string_pointer() {
	fmt.Println("pointer receiver OK！")
}


func main(){

	var s1 = IntSet{}
	s1.string() // ok
	var s2 = &IntSet{}
	s2.string() // ok

	var s3 = IntSet{}
	s3.string_pointer()

	var s4 = &IntSet{}
	s4.string_pointer()

	var s5 IntSet
	s5.string()
	s5.string_pointer() // ok

	var s6 = IntSet{}
    s6.string_pointer()

	//IntSet{}.string_pointer() // compile error
}