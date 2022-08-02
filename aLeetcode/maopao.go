package main

import "fmt"

func main()  {
	array := []int{5,4,3,4,2}
	res := bubleSort(array)
	fmt.Println(res)
}

func bubleSort(array []int) []int {
	length :=len(array)
	//isChange :=false

	for i:=0;i<length;i++ {
		for j:=0;j<length-i-1;j++ {
			// j <= length-i-1 这个是关键，每次 i ，少比较最后一位数组
			if array[j] > array[j+1] {
				array[j+1],array[j] = array[j],array[j+1]
				//isChange = true
			}
		}
		// 直接跳下次循环
		// if !isChange {
		//     break;
		// }
	}
	return array
}
