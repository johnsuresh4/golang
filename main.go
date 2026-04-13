package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site! </h1>")
}

func main() {
	http.HandleFunc("/", handleFunc)
	fmt.Println("Starting the server: 3000")
	http.ListenAndServe(":3000", nil)
}
