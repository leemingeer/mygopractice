package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main(){

	http.HandleFunc("/double", doubleHandler)
	log.Fatalln(http.ListenAndServe(":4000", nil))
}

// 第一个参数是接口，因为返回要求有固定的方法
// 第二个参数是struct，已经固定好了对象结构，包含方法，URL等信息
func doubleHandler(w http.ResponseWriter, r *http.Request){
	text := r.FormValue("v")
	if text == ""{
		//specified error message and HTTP code， 返回错误信息和错误码
		http.Error(w, "missing value", http.StatusBadRequest)
		return
	}
	v, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, "not a number: " + text, http.StatusBadRequest)
	}
	if _, err = fmt.Fprintln(w, v*2); err != nil{
		http.Error(w, "cannot write to response", http.StatusBadRequest)
		return
	}
}

// debug
//leemingeer@mbp go % curl "localhost:4000/double?v=3"
//6