package main

import (
"fmt"
"log"
"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	// w是接口类型，入参要求是io.Writer接口，但是w实现了io.Write的方法
	_, err := fmt.Fprintln(w, "Hello World!" )
	if err != nil{
		log.Fatalln(err)
	}

}


func main() {
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}