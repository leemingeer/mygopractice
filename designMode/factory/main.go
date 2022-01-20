
package main

import "strconv"

func evalRPN(tokens []string) int {

	stack := []int{}

	for _, token := range tokens {

		val, err := strconv.Atoi(token)

		if  err == nil{
			stack = append(stack, val)
		} else {
			stack[len(stack)-1] = calculate(stack[len(stack)-2], stack[len(stack)-1], token)
			stack = stack[:len(stack)-1]
		}
	}
	return stack[0]

}


type Addint1 int

func (a Addint1) calculate(num1, num2 int)int{
	return num1 + num2
}
type Minint int
func (a Minint) calculate(num1, num2 int)int{
	return num1 - num2
}

type Multiint int
func (a Multiint) calculate(num1, num2 int)int{
	return num1 * num2
}
type Divint int
func (a Divint) calculate(num1, num2 int)int{
	return num1 / num2
}
type I interface{
	calculate(num1, num2 int) int
}


func factory(oper string) I{

	var obj I

	if oper == "+"{
		obj = Addint1(2)
	}

	if oper == "*"{
        obj = Multiint(1)
	}
	if oper == "-"{
		obj = Minint(1)
	}
	if oper == "/"{
		obj =  Divint(1)
	}
	return obj
}

func calculate(num1, num2 int, oper string)int{
	i := factory(oper)
	return i.calculate(num1, num2)
}


