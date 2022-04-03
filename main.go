package main

import (
	"fmt"
	"time"
)

// I'm following a youtube tutorial to develop a basic booking application
// https://www.youtube.com/watch?v=yyUHQIec83I

var conferenceName = "Go conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

// Following can be used to declare on array of fixed size
// var bookings = [50]string{}
// An alternate declaration of an expanding/contracting array called slice
// Alternate slice declarations
// bookings := []string
// var bookings = []string{}
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	userEmail   string
	noOfTickets uint
}

func main() {
	greetUsers()

	// Loops introduction. We have just one type of loop in Go. For!
	// This is an infinite loop
	for {
		firstName, lastName, userEmail, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTickets := validateUserInput(firstName, lastName, userEmail, userTickets)

		if isValidName && isValidEmail && isValidUserTickets {
			bookTickets(userTickets, firstName, lastName, userEmail)
			sendTicket(userTickets, firstName, lastName, userEmail)
			firstNames := getFirstNames()
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

func greetUsers() {
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
}

func getFirstNames() []string {
	firstNames := []string{}
	// range returns index and the corresponding data from a slice
	// if your code doesn't require the index variable (you can use any name for your index and value of course)
	// you can instead use an underscore(_) so that the compiler doesn't complain about an unused variable
	// I think we are using strings.Fields to split booking data and get firstName here instead of the available
	// firstName variable, just so we can learn a new conecpt. And nothing more.
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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
	return firstName, lastName, userEmail, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, userEmail string) {
	remainingTickets = remainingTickets - userTickets
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		userEmail:   userEmail,
		noOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. Tickets will be sent to %v\n", firstName, lastName, userTickets, userEmail)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, userEmail string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("**************")
	fmt.Printf("Sending ticket:\n%v to email address %v\n", ticket, userEmail)
	fmt.Println("**************")
}
