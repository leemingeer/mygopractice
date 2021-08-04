package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request){
    fmt.Println("handler start")
	ctx := r.Context()

	anotherComplete := make(chan struct{})
	go func(){
		// do something
		time.Sleep(2 * time.Second)
		anotherComplete <- struct{}{}
	}()



	select { //监听通道
	case <- anotherComplete:
		fmt.Println("finished by innner goroutine after 2s")
	case<- time.After(5 * time.Second):
		fmt.Println("finish doing something after 5s")
	case <- ctx.Done(): // ctx is cancelled 提前结束
		err := ctx.Err()
		if err != nil{
			fmt.Println(err.Error())
		}
	}

	fmt.Println("handler ends")

}





func main() {
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(":8081", nil))
}