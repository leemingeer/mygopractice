package main

import "fmt"

func main() {
	i := 1
	fmt.Printf("%p\n", &i)
	ii := 1
	fmt.Printf("%p\n", &ii)
	i = 2
	fmt.Printf("%p\n", &i)

	//s1 := []string{}
	//s2 := []string{}
	// fmt.Println(s1 == s2) // can't compare two slice

	var strList []string
	var numList []int
	var numListEmpty = []int{}
	fmt.Printf("%p,%p,%p\n", &strList, &numList, &numListEmpty)
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	fmt.Println(strList == nil) // true, claim but not used, default nil
	fmt.Println(numList == nil) // true, claim but not used, default nil
	fmt.Println(numListEmpty == nil ) // false, already allocated memory, just has no element

    fmt.Println("----------")
	var a1 = []int{1,2,3}
	fmt.Printf("%p, %p\n", &a1, &a1[0])
	b :=a1
	fmt.Printf("%p\n",&b) // 问题1、为什么b的地址跟a不同，他们不是共享底层数组吗？
	fmt.Printf("%d, %d\n", a1, b)
	fmt.Printf("%p, %p\n", &a1[0], &b[0])
	b = append(b,0)
	//b[1]=999
	fmt.Printf("%p\n", &b) // 问题2、为什么b的地址维持不变？append时b的长度不够应该重新分配了内存空间吧
	fmt.Println(a1)
	fmt.Println(b)
	fmt.Printf("%d, %d\n", a1, b)
	fmt.Printf("%p, %p\n", &a1[0], &b[0])

	fmt.Println("-----slice-----")

	str1 := "hello world"
	str_slice := str1[1:8]
	fmt.Println(&str1, &str_slice)
    ss1 := []string{"helo", "world"}
    fmt.Println(ss1)
	ss2 := []byte{'h','e','l', 'o', ' ', 'w', 'o','r','l','d'}
	fmt.Println(ss2[1:8], string(ss2[1:8]))
	fmt.Println("----------")

	var a = []int{1,2,3}
	a = append([]int{0}, a...)
	fmt.Println(a)
	a = append(a, 4,5,6)
	fmt.Println(a)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	ret := copy(slice2, slice1)
	fmt.Println(slice1, slice2, ret)

	slice3 := []int{6, 7, 8}
	ret = copy(slice1, slice3)
	fmt.Println(slice1, slice2, slice3, ret)
}
