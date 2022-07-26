package main

import "fmt"

type Student struct{
	name string
	age int
}

type myInt int

func (m myInt) print(){
	fmt.Println(m)
}

func main(){
	var t1 []int
	fmt.Println(t1)
	t1 = append(t1 ,3 )
	t1[0] = 4
	fmt.Println(t1)

	var i myInt
	i.print()

}