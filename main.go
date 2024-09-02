package main

import (
	"fmt"
	"strings"
)

var bookings = []string{}
var conferanceName = "go conference"

const conferanceTicket = 50

var remainingTickets uint = 50

func main() {

	greetUsers()

	for {
		// get user inputs
		firstName, lastName, email, userTicket := getuserInput()

		//validate user inputs
		isvalidName, isvalidemail, isvalidTicketnumber := validateUserInput(firstName, lastName, email, userTicket, remainingTickets)

		// if all inputs are valid
		if isvalidName && isvalidemail && isvalidTicketnumber {

			bookTicket(userTicket, firstName, lastName, email)

			// call function to get first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings are : %v\n ", firstNames)

			// if ticket is not avalable then break
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {

			if !isvalidName {
				fmt.Println(" frist name or last name you entered is too short")

			}
			if !isvalidemail {
				fmt.Println("email address you entered is not valid")

			}
			if !isvalidTicketnumber {
				fmt.Println("number of tickets you entered is invalid")
			}

			// if  inputs are not valid
			fmt.Printf("your input data is invalid. try agin\n")
			continue

		}
	}

}

func greetUsers() {

	fmt.Printf("well come %v my frist application \n", conferanceName)
	fmt.Printf("we have total tikets %v and avalable %v tikets. \n ", conferanceTicket, remainingTickets)
	fmt.Println("book you ticket here")

}
func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {

		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func validateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool) {

	isvalidName := len(firstName) >= 2 && len(lastName) >= 2
	isvalidemail := strings.Contains(email, "@")
	isvalidTicketnumber := userTicket <= remainingTickets && userTicket > 0

	return isvalidName, isvalidemail, isvalidTicketnumber

}

func getuserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter Your First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your Last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your email:")
	fmt.Scan(&email)

	fmt.Println("amount of ticket ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket

}

func bookTicket(userTicket uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTicket
	//book ticket
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTicket, email)
	fmt.Printf("we have total tikets %v and avalable %v tikets. \n ", conferanceTicket, remainingTickets)

}
