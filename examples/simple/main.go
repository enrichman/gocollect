package main

import (
	"fmt"
	"net/http"

	_ "github.com/enrichman/gocollect"
)

func main() {
	http.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	}))

	fmt.Println(http.ListenAndServe(":8082", nil))
}
