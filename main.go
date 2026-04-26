package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string, user User) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, user)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// user := User{
	// 	Name: "John Suresh",
	// 	Age:  32,
	// 	Meta: UserMeta{
	// 		Visits: 8,
	// 	},
	// }
	// executeTemplate(w, "hello.gohtml", user)

	// bio := `&lt;script&gt;alert(&quot;Hi!&quot;);&lt;/script&gt;`
	w.Header().Set("content-type", "text/html; charset=utf-8")
	tplPath := filepath.Join("templates", "home.gohtml")
	tp1, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	err = tp1.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}

}

// fmt.Fprint(w, "<h1>Welcome to my awesome site! </h1>")

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
