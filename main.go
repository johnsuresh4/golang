package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site! </h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:johnsuresh777@gmail.com\">johnsuresh777@gmail.com</a></p>")
}

func noFoundHandler(w http.ResponseWriter, r *http.Request) {
	// http.Error(w, "Page not found", http.StatusNotFound)
	// simply way to write this.
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Page Not Found</h1>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path == "/" {
	// 	homeHandler(w, r)
	// 	return
	// } else if r.URL.Path == "/contact" {
	// 	contactHandler(w, r)
	// 	return
	// }

	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		// TODO: handle the page not found error
		noFoundHandler(w, r)
	}
}

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		// TODO: handle the page not found error
// 		noFoundHandler(w, r)
// 	}
// }

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)

	// http.Handler - interface with ServeHTTP Method

	// http.HandlerFunc - a function type that accepts same args as ServeHTTP method, also implements http.Handler

	// http.H

	fmt.Println("Starting the server: 3000")
	http.ListenAndServe(":3000", nil)

}
