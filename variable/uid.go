package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
)

func GenerateID() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func main(){
	fmt.Println(GenerateID()[:32], len(GenerateID()))
}