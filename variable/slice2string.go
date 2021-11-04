package main

import (
	"fmt"
	"strings"
)

func main(){
	var input = "pvc208a140c52954b3d945263d531a472dc"
	expect := "pvc-208a140c-5295-4b3d-9452-63d531a472dc"
	fmt.Println(recoverPVName(input) == expect)
}

func recoverPVName(pvName string) string{
	var pv []string
	for i,v := range(pvName){
		if i == 3 || i==11 || i==15 || i==19 ||i==23{
			pv = append(pv, "-")
		}
		pv = append(pv, string(v))
	}
	return strings.Join(pv, "")
}