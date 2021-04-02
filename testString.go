package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	tracer := "死神来了, 死神bye bye"
	comma := strings.Index(tracer, ", ")
	fmt.Println(comma, tracer[comma:])
	pos := strings.Index(tracer[comma:], "死神")
	fmt.Println(comma, pos, tracer[comma+pos:])


	angel := "Aeros never die"
	angleBytes := []byte(angel)
	fmt.Println(angleBytes)
	fmt.Println(string(angleBytes))
	for i := 5; i <= 10; i++ {
		angleBytes[i] = ' '
	}
	fmt.Println(string(angleBytes))


	hammer := "吃我一锤"
	sickle := "死吧"
	// 声明字节缓冲
	var stringBuilder bytes.Buffer
	// 把字符串写入缓冲
	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)
	// 将缓冲以字符串形式输出
	fmt.Println(stringBuilder.String())

	s1 := "lee"
	s2 := "ming"
	var build strings.Builder
	build.WriteString(s1)
	build.WriteString(s2)
	s3 := build.String()
	fmt.Println(s3)

}
