package main

import (
	"fmt"
	"io"
)

type Socket struct {
}

// SSocket 结构的 Write() 方法实现了 io.Writer 接口
func (s *Socket) Write(p []byte) (n int, err error) {
	fmt.Println("socket write is called")
	return 0, nil
}
func (s *Socket) Close() error {
	fmt.Println("socket close is called")
	return nil
}

// usingWriter() 和 usingCloser() 完全独立，互相不知道对方的存在，也不知道自己使用的接口是 Socket 实现的。
// 使用io.Writer的代码, 并不知道Socket和io.Closer的存在
func usingWriter(writer io.Writer) {
	writer.Write(nil) // 用接口变量去调，就看接口变量被谁实现了，所谓接口，谁实现了就对接谁
}

// 使用io.Closer, 并不知道Socket和io.Writer的存在
func usingCloser(closer io.Closer) {
	closer.Close()
}

func main() {

	// 实例化Socket
	s := new(Socket)
	// 通过形参方式，接口变量 = 接口实现的实例
	usingWriter(s)

	usingCloser(s)
}
