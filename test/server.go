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
	fmt.Println("Hello World")
}
