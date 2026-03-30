package main

import (
	"strings"
)

var MyVar string = "Hai Team"

func validateUserInput(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTicket := userTickets <= int(remainingTickets)

	return isValidEmail, isValidName, isValidUserTicket
}
