package main

import (
	"fmt"
	"time"
)
var timeNow = time.Now

func GetDate() int {
	return timeNow().Day()
}

func main(){
	//fmt.Println(time.Now())
	//
	//stubs := Stub(&timeNow, func() time.Time {
	//	return time.Date(2015, 6, 1, 0, 0, 0, 0, time.UTC)
	//})
	//defer stubs.Reset()
	//
	//fmt.Println(GetDate())
	fmt.Println(time.Now().Unix())
	time.Sleep(3 *time.Second)
	fmt.Println(time.Now().Unix())
}