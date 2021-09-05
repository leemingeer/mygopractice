package main

import (
	"fmt"
	"time"
)

func main(){
	arrays := []int{1,2,3,4}

	var finalarrays []int

	for _, v := range(arrays){
		finalarrays = append(finalarrays, v)
	}
	fmt.Println(finalarrays)


	var finalarrayspoint []*int
	for _, v := range(arrays){
		finalarrayspoint = append(finalarrayspoint, &v)
	}
	for _, v := range(finalarrayspoint){
		fmt.Println(*v)
	}

    // 容易出错的闭包问题
    // i 作为参数传入
	for i:=0; i<10; i++{

		go func(){
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)


}
