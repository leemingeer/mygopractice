package main

import "fmt"

func findunique( s string){

	for i:=0; i<len(s);i++{
		for j:=i+1;j<len(s);j++{

			if s[i] == s[j]{
				fmt.Println(string(s[i]), "==", string(s[j]))
				break
			}

		}
	}

}

func main(){
	s := "eetcode"
	//fmt.Println(s, s[0], 'a', s[1]-'a')
	findunique(s)

}
