package main

import (
	"fmt"
	"strconv"
	//"reflect"
)

func main(){
	maxVolumeSize := ""
	maxVolumeSizeGB, err := strconv.Atoi(maxVolumeSize)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(maxVolumeSizeGB)
}
