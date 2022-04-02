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
	var bookings [50]string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have totol of %v and we have %v available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var userEmail string
	var userTickets uint
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&userEmail)
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first booking: %v\n", bookings[0])

	fmt.Printf("Thank you %v %v for booking %v tickets. Tickets will be sent to %v\n", firstName, lastName, userTickets, userEmail)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}
