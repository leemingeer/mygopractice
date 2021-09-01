package main

import "fmt"

func test_string(s string){
	fmt.Printf("inner: %v, %v\n",s, &s)
	s = "b"
	fmt.Printf("inner: %v, %v\n",s, &s)
}

func test_map2(m map[string]string){
	fmt.Printf("inner: %v, %p\n",m, m)
	m = make(map[string]string, 0) // make申请新的内存
	fmt.Printf("inner after make: %v, %p\n",m, m)
	m["a"]="11"
	fmt.Printf("inner: %v, %p\n",m, m)
}

func test_chan2(ch chan string){
	fmt.Printf("inner: %v, %v\n",ch, len(ch))
	ch<-"b"
	fmt.Printf("inner: %v, %v\n",ch, len(ch))
}

func test_slice1(sl []string){
	fmt.Printf("%v, %p\n",sl, sl)
	sl[0] = "aa"
	//sl = append(sl, "d")
	fmt.Printf("%v, %p\n",sl, sl)
}

func test_slice2(sl []string){
	fmt.Printf("%v, %p\n",sl, sl)
	sl = append(sl, "d")
	fmt.Printf("%v, %p\n",sl, sl)
}



func main() {
	s := "a"
	fmt.Printf("outer: %v, %v\n",s, &s)
	test_string(s)
	fmt.Printf("outer: %v, %v\n",s, &s)

	var m map[string]string//未初始化

	fmt.Printf("outer: %v, %p\n",m, m)
	test_map2(m)
	fmt.Printf("outer: %v, %p\n",m, m)


	ch := make(chan string, 10)
	ch<- "a"
	fmt.Printf("outer: %v, %v\n",ch, len(ch))
	test_chan2(ch)
	fmt.Printf("outer: %v, %v\n",ch, len(ch))

	sl := []string{
		"a",
		"b",
		"c",
	}

	fmt.Printf("%v, %p\n",sl, sl)
	test_slice1(sl)
	fmt.Printf("%v, %p\n",sl, sl)

	sl = []string{
		"a",
		"b",
		"c",
	}
	fmt.Printf("%v, %p\n",sl, sl)
	test_slice2(sl)
	fmt.Printf("%v, %p\n",sl, sl)
}