package main

import (
	"fmt"
	"strings"
)

// I'm following a youtube tutorial to develop a basic booking application
// https://www.youtube.com/watch?v=yyUHQIec83I

func main() {
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// Following can be used to declare on array of fixed size
	// var bookings = [50]string{}
	// An alternate declaration of an expanding/contracting array called slice
	var bookings []string
	// Alternate slice declarations
	// bookings := []string
	// var bookings = []string{}

	// %v format specifier is a generic one that can be used with multiple data types.
	// %v prints values, %+v prints the field and value (think strcutures)
	// %#v prints the struct along with field and value
	// Println adds new line in the end. Also adds space in between it's arguments
	// For example Println("hello", "world") -> "hello world"
	// Printf same as Println, except it doesn't add new line in the end
	// and space is added in between the arguments if neither of them are strings
	// Print is same as Printf
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have totol of %v tickets and we have %v tickets available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// Loops introduction. We have just one type of loop in Go. For!
	// This is an infinite loop
	for {
		var firstName string
		var lastName string
		var userEmail string
		var userTickets uint

		// Scan can be used to get input from the user
		// & is an address operator
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&userEmail)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) > 1 && len(lastName) > 1
		isValidEmail := strings.Contains(userEmail, "@")
		isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidUserTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. Tickets will be sent to %v\n", firstName, lastName, userTickets, userEmail)
			fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			// range returns index and the corresponding data from a slice
			// if your code doesn't require the index variable (you can use any name for your index and value of course)
			// you can instead use an underscore(_) so that the compiler doesn't complain about an unused variable
			// I think we are using strings.Fields to split booking data and get firstName here instead of the available
			// firstName variable, just so we can learn a new conecpt. And nothing more.
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("The %v is sold out\n", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email you entered doesn't contain @")
			}
			if !isValidUserTickets {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
}

// Understanding switch 
city := "London"

switch city {
	case "New York":
		// code for NY
	case "Singapore", "London":
		// code for Singapore or London. Multiple case consolidation
	case "Berlin":
		// code for Berlin
	case "Mexico City", "Hong Kong":
		// code for Mexico City or Hong Kong. Multiple cases
	default:
		fmt.Print("No valid city selected")
}