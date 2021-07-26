package race

import (
	"fmt"
	"sync"
	"sync/atomic"
)


func unsafeIncCounter() {
	defer wg.Done()
	for i:=0; i < 10000; i++{
		// race condition here
		counter ++
		//temp := counter  读
		//temp = temp + 1  运算
		//counter = temp   写入

	}

}


func MutexIncCounter() {
	defer wg.Done()
	for i:=0; i < 10000; i++{
		mtx.Lock()
		counter ++
		mtx.Unlock()
	}
}

// 使用channel进行同步
func ChannelIncCounter(){
	defer wg.Done()
	for i:=0; i < 5; i++{
        count := <- ch
        fmt.Printf("get count: %d\n", count)
        count ++
        ch <- count
	}



}

func AtomicIncCounter() {
	defer wg.Done()
	for i:=0; i < 10000; i++{
		atomic.AddInt32(&counter, 1)
	}
}





var wg sync.WaitGroup
var counter int32

var mtx sync.Mutex

var ch = make(chan int, 1)

func Mainrace() {


	wg.Add(2)
    //go MutexIncCounter()
	//go MutexIncCounter()
	//go AtomicIncCounter()
	//go AtomicIncCounter()

	go ChannelIncCounter()
	go ChannelIncCounter()
	ch <- 10 // 丢进去一个数，最后再拿出来
	wg.Wait() //直到两个player推出, 等待两个goroutine退出，channel 声明是有缓冲的，否则最后一个goroutine在往channel写入数据后，一直没有goroutine取出数据，导致卡在这里不退出，go wait产生死锁。
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
