package main

import (
	"fmt"
	"strings"
	"booking-app/helper"
)
//variable outside of all function are called PACKAGE LEVEL VARIABLES

var conferenceName = "Go conference" // conferenceName :="Go conference" // means the same thing and it doesn't work on the constants and variable that are explicitly
const conferenceTickets uint = 50
var remainingTickets uint = 50
var bookings [50]string
var bookingSlice = []string{}


func main() {

	// Another function call
	greetUser()
  
	// using lOOP
	for {

		firstName, lastName, email, userTicketsCount := getUserInput()
		// validations
		isValidName, isValidEmail, isValidTicket := helper.ValidateUserInputs(firstName, lastName, email, userTicketsCount, remainingTickets)
        
		// conditions
		if isValidName && isValidEmail && isValidTicket {
			bookTicket(userTicketsCount, firstName, lastName, email)
			
			

			// Array
			fmt.Printf("The whole array: %v\n", bookings)
			fmt.Printf("the first value: %v \n", bookings[0])
			fmt.Printf("Array type %T \n", bookings)
			fmt.Printf("The Array length: %v \n", len(bookings))

			// Slice
			fmt.Printf("The whole slice: %v\n", bookingSlice)
			fmt.Printf("the first value: %v \n", bookingSlice[0])
			fmt.Printf("Slice type %T \n", bookingSlice)
			fmt.Printf("The Slice length: %v \n", len(bookingSlice))

			// call function print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of booking are: %v \n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is sold out please come next year")
				break
			}

		} else if userTicketsCount == remainingTickets {

			fmt.Printf(" you are buying every ticket %v tickets remaining, so you can't book %v tickets \n", remainingTickets, userTicketsCount)

		} else {

			if !isValidEmail {
				fmt.Printf("Your email are invalid")
			} else if !isValidName {
				fmt.Printf("Either of your names is not invalid")
			} else if !isValidTicket {
				fmt.Printf("Invalid number ofntickers ")
			}
		}
	}

}
func greetUser() {
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "ticket and ", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	// print formatting
	fmt.Printf("Welcome to %v booking application \n remaining ticket %v \n ", conferenceName, remainingTickets)

}

// function which returns a slice of string
func getFirstNames() []string {
	firstNames := []string{}
	for _, aBooking := range bookingSlice {
		var names = strings.Fields(aBooking)
		firstNames = append(firstNames, names[0])

	}
	return firstNames
}



func getUserInput() (string, string, string, uint) {
	// address
	var firstName string
	var lastName string
	var email string
	var userTicketsCount uint //accept only positive integers

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address : ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets : ")
	fmt.Scan(&userTicketsCount)

	return firstName,
		lastName,
		email,
		userTicketsCount

}
func bookTicket(userTicketsCount uint, firstName string, lastName string, email string){
	remainingTickets = remainingTickets - userTicketsCount
	bookings[0] = firstName + " " + lastName
	bookingSlice = append(bookingSlice, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets, You will receive a confirmation email at %v \n", firstName, lastName, userTicketsCount, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}