package main

import "fmt"

func f1() (result int) {
	result = 0
	//在return之前，执行defer函数
	defer func() {
		result++
	}()
	return
}

func f2() (r []int) {
	t := []int{5}
	//赋值

	r  = t
	//在return之前，执行defer函数，defer函数没有对返回值r进行修改，只是修改了变量t
	defer func() {
		t[0] = t[0] + 5
	}()
	return
}

func f3() (r int) {
	//给返回值赋值
	r = 1
	/**
	 * 这里修改的r是函数形参的值，是外部传进来的
	 * func(r int){}里边r的作用域只该func内，修改该值不会改变func外的r值
	 */
	defer func(r int) {
		r = r + 5
	}(r)
	return
}

func main(){
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}