package main

import (
	"fmt"
	"go-booking-app/input"
)

func getUserDetailInput() (string, string, string) {

	firstName := input.GetString("Please, insert your first name")
	lastName := input.GetString("Please, insert your last name")
	email := input.GetString("Please, insert your email")

	return firstName, lastName, email
}

func getTicketAmountInput() uint {
	prompt := fmt.Sprintf("Please, type in how many tickets you want to buy (max %d)", remainingTickets)
	userTickets := input.GetInt(prompt)
	return uint(userTickets)
}
