package main

import "fmt"
type Level int8
const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel

	_minLevel = DebugLevel
	_maxLevel = FatalLevel
)

const (
	a1 = iota - 1   // a1 = 0   // 又一个const出现, iota初始化为0
	a2 = iota   // a1 = 1   // const新增一行, iota 加1
	a3 = 6      // a3 = 6   // 自定义一个常量
	a4          // a4 = 6   // 不赋值就和上一行相同
	a5 = iota   // a5 = 4   // const已经新增了4行, 所以这里是4
)

const (
	a6 = iota     //b=0
	a7            //c=1   相当于c=iota
	a8
)

const (
	aa = iota - 1   // -1
	bb              // 0
	cc              // 1
	dd = 2
	ee = iota      //
)
func main(){
	fmt.Println(DebugLevel, ErrorLevel)
	fmt.Println(a1,a2,a3,a4,a5,a6,a7,a8)
	fmt.Println(aa, bb, cc,dd,ee)
}