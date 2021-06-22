package main

import "fmt"

func nonempty(strings []string) []string{
	i := 0
	for _, s := range strings{
		if s != ""{
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main(){
	data := []string{"one", "", "three"}
	fmt.Printf("origin: %q\n", data)
	fmt.Printf("nonempty: %q\n", nonempty(data))
	fmt.Printf("after: %q\n", data)

}
