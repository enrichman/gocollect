package main

import (
	"fmt"
	"net/http"

	"github.com/enrichman/gocollect"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	}))

	mux.HandleFunc("/collect", gocollect.Collect("tmp"))

	fmt.Println(http.ListenAndServe(":8082", mux))
}
