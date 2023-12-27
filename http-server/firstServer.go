package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Req received")
		fmt.Println(r.Method)
		w.Write([]byte("Hello World"))
	})
	mux.HandleFunc("/getgoing/go", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Req received")
		fmt.Println(r.Method)
		w.Write([]byte("Hello World go"))
	})
	http.ListenAndServe("localhost:3000", mux)
}
