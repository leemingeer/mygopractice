package main

import "fmt"


type Mover2 interface {
	move()
}

type dog2 struct {}

func (d *dog2) move() {
	fmt.Println("狗会动")
}
func main() {
	var x2 Mover2
	//var wangcai = dog2{} // 旺财是dog类型
	//x2 = wangcai         // x不可以接收dog类型
	var fugui = &dog2{}  // 富贵是*dog类型
	x2 = fugui           // x可以接收*dog类型
	x2.move()
}