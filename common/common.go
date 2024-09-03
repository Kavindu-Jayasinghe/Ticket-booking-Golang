package common

import "strings"

// ValidateUserInput checks the validity of user inputs.
// It returns three booleans indicating if the name, email, and ticket number are valid.
func ValidateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket <= remainingTickets && userTicket > 0

	return isValidName, isValidEmail, isValidTicketNumber
}
