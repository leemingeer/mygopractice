package main

import "fmt"

func reverse(s []string) []string{
	// 不支持i++,j--这种操作，原因是逗号不能连接多个语句，但是可以有多个赋值
	// for i,j :=0, len(s)-1; i<j; i++,j--{
    // swap
	//}

	for i,j :=0, len(s)-1; i<j; i,j = i+1,j-1{
		s[i], s[j] = s[j], s[i]
	}
	return s
}


func main(){
	fmt.Println(5/2) // 2
	s1 := []string{"I", "am", "xiao", "ming",  "o"}
	s2 := []string{"I", "am", "xiao", "ming"}
	// 就地翻转slice、字符串
	fmt.Printf("origin: %s\n", s1)
	fmt.Printf("reverse: %s\n", reverse(s1))
	fmt.Printf("origin: %s\n", s2)
	fmt.Printf("reverse: %s\n", reverse(s2))

	s1 = []string{"I", "am", "xiao", "ming",  "o"}
    // [xiao ming o I am]
	fmt.Println("向左移动两位")
	reverse(s1[:2])
	reverse(s1[2:])
	reverse(s1[:])
	fmt.Println(s1)

	fmt.Println("向右移动两位")
	//[ming o I am xiao]
	s1 = []string{"I", "am", "xiao", "ming",  "o"}
	reverse(s1[len(s1) -2:])
	reverse(s1[:len(s1) -2])
	reverse(s1[:])
	fmt.Println(s1)



	
}
