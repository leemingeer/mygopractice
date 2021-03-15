package main

import "fmt"

var s []int
// s = append(s, 1)
// go和c一样，所有的运算都应该在函数内进行，函数外进行的是语法错误。
// 函数体外初始化结构体就两个办法，要么一次性全部赋值，要么先声明（全局/局部）变量，在某个函数内进行赋值，在函数体外进行结构体成员赋值相当于函数外面进行运算了。

// go不支持在函数外赋值与调用操作, 但可以声明一个变量或者声明并初始化一个变量。

func main(){
	s = append(s, 1)
	fmt.Print(s)
	//[1]
    f2()
	f3()
	f4()
}



type person struct {
	name string
	age int
}

// var P person
//P.name = "annie"
//P.age = 20

var P1 = person{"annie", 20}

func f2() {
	fmt.Printf("The person's name is %s\n", P1.name)
}


var P2 person

P2.name = "a"
P2.age = 10
fmt.Printf("The person's name is %s\n", P2.name)

}

func f4(){
	var s []int
	s = make([]int, 0)
	s = append(s, 1)
	fmt.Printf("%v\n", s)
}