package main

import (
	"booking-app/common"
	"fmt"
	"strings"
)

var bookings = []string{}
var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

func main() {

	greetUsers()

	for {
		// get user inputs
		firstName, lastName, email, userTicket := getUserInput()

		// validate user inputs
		isValidName, isValidEmail, isValidTicketNumber := common.ValidateUserInput(firstName, lastName, email, userTicket, remainingTickets)

		// if all inputs are valid
		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTicket, firstName, lastName, email)

			// call function to get first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			// if ticket is not available, then break
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {

			if !isValidName {
				fmt.Println("The first name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered is not valid.")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid.")
			}

			// if inputs are not valid
			fmt.Println("Your input data is invalid. Please try again.")
			continue

		}
	}

}

func greetUsers() {

	fmt.Printf("Welcome to %v, my first application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Book your ticket here:")

}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {

		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter Your First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your Last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your Email:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets:")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket

}

func bookTicket(userTicket uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTicket
	// book ticket
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

}
