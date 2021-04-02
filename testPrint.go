package main

import "fmt"

func main() {
	pi := 3.14159
	// 按数值本身的格式输出
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)
	fmt.Println(variant)
	// 匿名结构体声明, 并赋予初值
	profile := &struct {
		Name string
		HP   int
	}{
		Name: "rat",
		HP:   150,
	}
	fmt.Printf("使用'%%v' %v\n", profile) //&{rat 150}
	fmt.Printf("使用'%%v' %v\n", *profile) //{rat 150}
	fmt.Printf("使用'%%+v' %+v\n", profile) // &{Name:rat HP:150}
	fmt.Printf("使用'%%#v' %#v\n", profile) // &struct { Name string; HP int }{Name:"rat", HP:150}
	fmt.Printf("使用'%%T' %T\n", profile) // *struct { Name string; HP int }
	
}
