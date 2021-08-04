package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main(){

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3 * time.Second)
	defer cancel() // 用户ctrl+C停掉
	// 相比于单进程，如何将ctx传给server?
    // 需要新建请求
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8081", nil)
	if err != nil {
		log.Fatal(err)
	}
	req = req.WithContext(ctx) // 将r的context改成传入的ctx， 并且返回一个r的浅拷贝，从而将新ctx将入到req中

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//resp, err := http.Get("http://localhost:8080")
	//if err != nil{
	//	log.Fatalln(err)
	//}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(strings.Trim(string(respBytes), "\n"))


}
