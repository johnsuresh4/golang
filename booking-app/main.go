package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	var conferenceTickets int = 50
	var remainingTickets uint = 50
	// var bookings [50]string // This is array
	// var bookings = []string{} // This is slice
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their names
	fmt.Print("Please enter your First name: ")
	fmt.Scan(&firstName)

	fmt.Print("Please enter your Last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Please enter your Email: ")
	fmt.Scan(&email)

	fmt.Print("Please enter number of ticket: ")
	fmt.Scan(&userTickets)

	remainingTickets = uint(conferenceTickets) - uint(userTickets)

	fmt.Printf("%v %v has booked %v tickets for %v\n", firstName, lastName, userTickets, conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Println(bookings)
}
