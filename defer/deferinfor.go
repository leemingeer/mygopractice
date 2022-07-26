package main

import "fmt"

func main(){

	var ret int
	s := []int{3,6,9}

		for _, v := range s {
			func() {
				defer func() {
					v++
					ret = v
					fmt.Println("defer: ", ret)
				}()
				ret = v
				fmt.Println(ret)
			}()
	}
}
