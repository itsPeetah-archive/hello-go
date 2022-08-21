package main

import (
	"fmt"
	"sync"
	"time"
)

// Custom struct type
type BookingDetails struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// PACKAGE LEVEL variables/constants

// Go has type inference
const conferenceName = "Go Conference"
const maxConferenceTickets uint = 50 // override type inference
var remainingTickets = maxConferenceTickets

// array (size definition, either hard coded value or constant)
// var bookings [maxConferenceTickets]string
// slice (no size definition) (slice == dynamic list) (empty assignment)
// var bookings = []string{}

// making a slice of maps or structs
// var bookings = make([]map[string]string, 0) // initial size = 0
var bookings = make([]BookingDetails, 0) // initial size = 0

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// In Go we only have for loops (substitutes all kinds of loops)
	/*
		infinite loop:
			for {
				...do stuff...
			}
		same as:
			for true {
				...do stuff...
			}
	*/

	// While loop
	for remainingTickets > 0 && len(bookings) < 50 {
		// No type inference if variables are not initialized, we need to declare the type

		firstName, lastName, email := getUserDetailInput()

		// Input validation
		if !validateUserDetails(firstName, lastName, email) {
			fmt.Println("Invalid input: please try again.")
			continue
		}

		userTickets := getTicketAmountInput() // syntactic sugar for variables (type inference, not for constants)
		isOverbooking, invalidTicketAmount, bookingAllTickets := validateTicketAmount(userTickets)

		// Continue
		if isOverbooking {
			fmt.Printf("We only have %d tickets remaining!\n", remainingTickets)
			continue
		} else if invalidTicketAmount {
			fmt.Println("Please insert a valid number of tickets...")
			continue
		} else if bookingAllTickets {
			fmt.Println("Wow, you're buying all the remaining tickets!")
		} else {
			fmt.Println("Proceeding to booking your tickets...")
		}

		details := bookTickets(userTickets, firstName, lastName, email)

		printBookingConfirmation(firstName, lastName, email, userTickets)
		printBookingFirstNames()

		wg.Add(1) // Add 1 new thread to wait for

		// Go starts a new goroutine (lightweight thread)
		go sendTicket(details)

		// Break, if-else
		if remainingTickets < 1 {
			fmt.Println("Our conference is fully booked up!")
			break
		}
	}

	// Wait for all the threads in the waitgroup before exiting
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %s booking application.\n", conferenceName)
	fmt.Printf("We have %d out of %d tickets remaining.\n", remainingTickets, maxConferenceTickets)
	fmt.Println("Get your tickets here to attend")
}

func bookTickets(tickets2Book uint, firstName, lastName, email string) BookingDetails {

	// Create map for user
	// Maps can only have a single data type for keys and a single data type for values
	// userDetails := make(map[string]string)
	// userDetails["lName"] = lastName
	// userDetails["email"] = email
	// userDetails["numTcks"] = strconv.FormatUint(uint64(tickets2Book), 10)

	// let's use a struct
	userDetails := BookingDetails{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: tickets2Book,
	}

	remainingTickets -= tickets2Book
	bookings = append(bookings, userDetails) // adding a value to a slice

	return userDetails
}

func printBookingFirstNames() {
	// For each loop
	var bookingFirstName string
	firstNamesInBookings := []string{}
	for _, booking := range bookings {
		// bookingName = strings.Split(booking, " ")[0]
		// bookingFirstName = strings.Fields(booking)[0] // auto splits on emty spaces
		bookingFirstName = booking.firstName
		firstNamesInBookings = append(firstNamesInBookings, bookingFirstName)
	}
	fmt.Println("These are all the bookings in the application", firstNamesInBookings)
	fmt.Println("Remaining tickets:", remainingTickets)
}

// Same type argument declared in a row
func printBookingConfirmation(fName, lName, email string, tickets uint) {
	fmt.Printf("Your details:\n - Name: %s %s\n - Email: %s\n", fName, lName, email)
	fmt.Printf("Thank you for booking %d tickets!\n", tickets)
}

// GO-ROUTINES

func sendTicket(dets BookingDetails) {
	fmt.Println("Sending tickets...")
	time.Sleep(10 * time.Second) // Simulate delay
	var tickets = fmt.Sprintf("%v tickets to %v %v.", dets.numberOfTickets, dets.firstName, dets.lastName)
	fmt.Println("###############")
	fmt.Printf("Sent ticket\n\t[%v]\nto %v.\n", tickets, dets.email)
	fmt.Println("###############")

	// removes the thread from the waiting list
	wg.Done()
}
