package main

import (
"fmt"
"sync"
"sync/atomic"
"time"
)

var G_Int int64;
var WG sync.WaitGroup;
var ThreadCnt int

func AtmoicOpr()  {
	var TempInt int64
	for {
		//等待所有协程启动完成，模拟并发问题
		if ThreadCnt == 100 {
			break;
		}
	}
	//错误写法
	//TempInt = G_Int
	//Result := atomic.CompareAndSwapInt64(&G_Int, TempInt, TempInt + 1)
	//fmt.Println(TempInt, " Try to CAS: " ,Result)
	//正确写法
	for {
		TempInt = atomic.LoadInt64(&G_Int)
		Result := atomic.CompareAndSwapInt64(&G_Int, TempInt, TempInt + 1)
		if Result == true {
			fmt.Println(TempInt, " Try to CAS: " ,Result)
			break;
		}
	}
	WG.Done()
}

func main() {
	G_Int = 0
	ThreadCnt = 0
	for i:= 0; i< 100; i++ {
		go AtmoicOpr()
		WG.Add(1)
		ThreadCnt += 1
		fmt.Println("ThreadCnt is: ", ThreadCnt)
	}
	WG.Wait()
	time.Sleep(time.Second * 2)
}