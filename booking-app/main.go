package main

import (
	"fmt"
	"strconv"
)

var conferenceName string = "Go Conference"
var conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

// Above statement, we are creating list of maps

// var bookings [50]string // This is array
// var bookings = []string{} // This is slice
// bookings := []string{} //This syntax can be used in func body only

func main() {
	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserData()

		isValidEmail, isValidName, isValidUserTicket := validateUserInput(firstName, lastName, email, userTickets)

		if !isValidName {
			fmt.Println("Please enter more than 3 chars")
			continue
		}

		if !isValidEmail {
			fmt.Println("Please enter valid email")
			continue
		}

		if isValidUserTicket {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		remainingTickets = uint(remainingTickets) - uint(userTickets)

		// create a map for user
		var userData = make(map[string]string)
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["numberOfTickets"] = strconv.FormatInt(int64(userTickets), 10)

		bookings = append(bookings, userData)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		printFirstNames()

		var noTicketsRemaining bool = remainingTickets == 0

		if noTicketsRemaining {
			// end the program
			fmt.Println("Our conference is booked out. Come next year.")
			break
		}

	}

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func getUserData() (string, string, string, int) {
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

	return firstName, lastName, email, userTickets
}

func printFirstNames() {
	firstNames := []string{}

	for _, booking := range bookings {
		// var firstName = names[0]
		firstNames = append(firstNames, booking["firstName"])
	}
	fmt.Printf("The first names of bookings are: %v\n", firstNames)

}
