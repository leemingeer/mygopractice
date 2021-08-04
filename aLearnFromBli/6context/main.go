package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context){

	select { // 监听ctx是否被关闭
	case<- time.After(5 * time.Second): // 相当于sleep(5s), 处理完了，返回只读通道，当设置的d时间过去后，会在通道传入一个当前的时间，因此监听这个通道就能得知是否已经过了d个单位的时间
		fmt.Println("finish doing something")
	case <- ctx.Done(): // ctx is cancelled 提前结束
	    err := ctx.Err()
	    if err != nil{
	    	fmt.Println(err.Error())
		}
	}
	time.Sleep(5 * time.Second)
	fmt.Println("finish doing something")
}

// 单进程
func main() {
	// 创建空context的两种方法, Context是接口类型
	ctx := context.Background() // 根context， 空context, 没有deadline， 不能被取消cancel, kv为空
	//todoCtx := context.TODO() //当不确定context的类型时使用

	//派生context
	// 利用空context（root context）来派生出新的context， 方法有三种
	ctx, cancel := context.WithCancel(ctx) //常用
	//ctx, cancel := context.WithTimeout(ctx, 2 time.Seconds) // 两秒后会自动停止
	//ctx, cancel := context.WithDeadline(ctx, 时间节点) // 具体的时间点到了，停止
	// 每个都返回一个新的context和一个cancel函数，我们可以调用cancel()
	//cancel 是func类型，当想取消这个ctx时，就调用这个cancel() func， 所有拿到这个ctx的goroutine会知道cancel了
	go func(){
		time.Sleep(6 * time.Second)
		cancel() // ctx没有被使用，不会通知到它，监听Context是否被取消， Context.Done()返回一个只读的通道，当context被取消时，这个通道里会传入一个值
		// 表示它已经被关闭了，因此我们只需要用select来监听context.Done()是否有值，就可以判断context是否被取消。
	}()


	doSomething(ctx)

}