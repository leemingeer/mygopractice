package main

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func main(){
	name := uuid.New().String()
	fmt.Println(name)
	fmt.Printf("%#v", []byte("abc"))
	fmt.Println(strings.Split(name,"-")[0])
	fmt.Println(strings.Replace(name,"-", "", -1))

	parameters := map[string]string{}
	parameters["portals"]= "[\"192.68.201.5:80\"]"
	fmt.Println(parameters)
}