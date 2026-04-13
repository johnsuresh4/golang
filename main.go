package main

import "fmt"

type Person struct {
	name string
	age  int
	f    func(string)
}

func getName(p Person) string {
	return p.name
}

// func handleFunc(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "<h1>Welcome to my awesome site! </h1>")
// }

func main() {
	p1 := Person{age: 23}
	p1.name = "Suresh"
	p1.age = 34
	p1.f = func(s string) {
		fmt.Println(s + "s")
	}
	p1.f("Hello")
	p2 := Person{age: 24}
	p2.name = "Suresh"
	p2.age = 34
	p2.f = func(s string) {
		fmt.Println(s + "asd")
	}
	p2.f("Hello")
}

// func helloUser(w http.ResponseWriter, r *http.Request) {
// 	var greeting = "Hello user. Welcome to our Todolist Application"
// 	fmt.Fprintln(w, greeting)
// }
