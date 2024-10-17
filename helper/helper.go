package helper

import "strings"

func ValidateUserInputs(firstName string, lastName string, email string, userTicketsCount uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTicketsCount > 0 && userTicketsCount <= remainingTickets

	return isValidName, isValidEmail, isValidTicket // other programmig language return just one value but go can return many values
}