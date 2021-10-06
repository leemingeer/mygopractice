package main

import "fmt"

type Myint int

func (n Myint) show(){
	fmt.Println("Hello World")
}

type Mystring string

func (n Mystring) show(){
	fmt.Println("Hello World")
}

type Mystring2 Mystring

func (n Mystring2) show2(){
	fmt.Println("Hello World")
}

func main(){

	c := Myint('c')
	c.show()

	s := Mystring("c")
	s.show()

	s2 := Mystring2("c")
	s2.show2()

}
