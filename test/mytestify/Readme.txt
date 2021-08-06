才会在本地目录编译出binary

case1: main包里没有main function
$ go build charge.go
# command-line-arguments
runtime.main_main·f: function main is undeclared in the main package

case2: main function的包名不是package main

编译不报错，但是本地目录没有可执行文件

case3: func main() + package main
才会在本地目录编译出binary