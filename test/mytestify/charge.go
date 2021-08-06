package mytestify

import (
	"fmt"
)

var A =  func()string{ return "hello world"}()

// MessageService 通知客户被收取的费用
type MessageService interface {
	SendChargeNotification(int) bool
}
// SMSService 是 MessageService 的实现
type SMSService struct{}
// SendChargeNotification 通过 SMS 来告知客户他们被收取费用
// 这就是我们将要模拟的方法
func (sms SMSService) SendChargeNotification(value int) bool {
	fmt.Println("Sending Production Charge Notification")
	return true
}


// MyService 使用 MessageService 来通知客户
type MyService struct {
	messageService MessageService  // field是接口类型变量，需要给他赋值
}
// ChargeCustomer 向客户收取费用
// 在真实系统中，我们会模拟这个
// 但是在这里，我想在每次运行测试时都赚点钱
func (a MyService) ChargeCustomer(value int) bool {
	a.messageService.SendChargeNotification(value)
	fmt.Printf("Charging Customer For the value of %d\n", value)
	return true
}

// 一个 "Production" 例子
func main() {
	fmt.Println("Hello World")

	smsService := SMSService{}
	myService := MyService{smsService} // 结构体给接口变量赋值
	myService.ChargeCustomer(100)
}