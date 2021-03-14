package main

import (
	"fmt"
)

func main(){
	// m1 原生普通方式
	var mapp map[string]float32
	mapp = make(map[string]float32)
	mapp["name"] = 11.11
	fmt.Printf("%v\n", mapp)

	//m2 短声明
	mappp := make(map[string]float32)
	mappp["sex"]=22.22
	fmt.Printf("%v\n", mappp)
	fmt.Println(mappp)

	//m3 结合
	var mapppp = make(map[string]float32)
	mapppp["school"] = 33.33
	fmt.Printf("%v\n", mapppp)
	fmt.Println(mapppp)
}