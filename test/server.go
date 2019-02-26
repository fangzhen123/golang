package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe("127.0.0.1:9230", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	var arr [10]int
	count := 0
	for i:=0;i<10 ;i++  {
		arr[i] = 1
	}
	for _,v := range arr {
		count = count + v
	}
	fmt.Println(count)
}
