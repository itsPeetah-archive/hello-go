package main

import "strings"

// No need to declare functions in order of execution
func validateUserDetails(fName string, lName string, email string) bool {
	return validateName(fName, lName) && validateEmail(email)
}

func validateName(fName string, lName string) bool {
	return len(fName) >= 2 && len(lName) >= 2
}

func validateEmail(email string) bool {
	return strings.Contains(email, "@")
}

// returning multiple values (can actually be different types)
func validateTicketAmount(desiredTickets uint) (bool, bool, bool) {
	overbooking := desiredTickets > remainingTickets
	invalidAmount := desiredTickets <= 0
	allTickets := desiredTickets == remainingTickets
	return overbooking, invalidAmount, allTickets
}
