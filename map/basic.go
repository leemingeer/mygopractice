package main

import (
	"fmt"
)

func main() {

	var ages map[string]int                   // 声明但没有初始化
	fmt.Printf("ages: %v, %p\n", ages, &ages) // map[], 0xc00000e038
	fmt.Println(ages == nil)                  // "true"
	fmt.Println(len(ages) == 0)               // "true"
	//ages["carol"] = 21  // panic: assignment to entry in nil map

	ages2 := map[string]int{}                    // 声明但初始化
	fmt.Printf("ages2: %v, %p\n", ages2, &ages2) //map[], 0xc00000e040
	fmt.Println(ages2 == nil)                    // "false", allocate memory
	fmt.Println(len(ages2) == 0)                 // "true"
	ages2["carol"] = 21
	fmt.Printf("%v", ages2) //map[carol:21]

	mapp := make(map[string]float32)
	fmt.Println(mapp == nil)           // false,
	fmt.Printf("%v,%p\n", mapp, &mapp) //map[], 0xc000006028
	mapp["name"] = 11.11
	mapp["age"] = 32
	fmt.Printf("%v, %f, %f\n", mapp, mapp["age"], mapp["location"])
}
