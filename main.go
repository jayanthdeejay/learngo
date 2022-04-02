package main

import (
	"fmt"
)

// I'm following a youtube tutorial to develop a basic booking application
// https://www.youtube.com/watch?v=yyUHQIec83I

func main() {
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// Following can be used to declare on array
	// var bookings = [50]string{}
	// An alternate declaration
	var bookings []string
	// Alternate slice declarations
	// bookings := []string
	// var bookings = []string{}

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
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&userEmail)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. Tickets will be sent to %v\n", firstName, lastName, userTickets, userEmail)
		fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

		fmt.Printf("Bookings list: %v\n", bookings)
	}
}
