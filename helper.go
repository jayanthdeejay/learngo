package main

import "strings"

// We need to capitalize the funtion name for us to be able to export this function and use it in another package
func validateUserInput(firstName string, lastName string, userEmail string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 1 && len(lastName) > 1
	isValidEmail := strings.Contains(userEmail, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidUserTickets
}
