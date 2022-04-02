package main

import (
	"fmt"
)

// I'm following a youtube tutorial to develop a basic booking application
// https://www.youtube.com/watch?v=yyUHQIec83I

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have totol of %v and we have %v available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}
