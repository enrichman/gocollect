package main

import (
	"fmt"
	"net/http"

	_ "github.com/enrichman/gocollect"
)

func main() {
	mux := http.DefaultServeMux // YES!

	mux.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	}))

	fmt.Println(http.ListenAndServe(":8082", mux))
}
