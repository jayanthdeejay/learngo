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

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have totol of %v and we have %v available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	fmt.Println(remainingTickets)
	fmt.Println(&remainingTickets)

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

}
