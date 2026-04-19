package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("hello.gohtml")
	w.Header().Set("content-type", "text/html; charset=utf-8")

	if err != nil {
		panic(err)
	}

	user := User{
		Name: "John Suresh",
		Age:  32,
		Meta: UserMeta{
			Visits: 8,
		},
	}

	error := t.Execute(w, user)

	if error != nil {
		panic(error)
	}
	// fmt.Fprint(w, "<h1>Welcome to my awesome site! </h1>")
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

type User struct {
	Name string
	Age  int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.NotFound(noFoundHandler)

	fmt.Println("Starting the server: 3000")
	http.ListenAndServe(":3000", r)

}
