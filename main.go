package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

//variable outside of all function are called PACKAGE LEVEL VARIABLES

var conferenceName = "Go conference" // conferenceName :="Go conference" // means the same thing and it doesn't work on the constants and variable that are explicitly
const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings [50]string
var bookingSlice = make([]UserData, 0) // a slice of structure

// structure: saves in ke-value pair in oder to hold values of different data-types
// Struck can be compared to classes in Java and other OOP languages
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// Another function call
	greetUser()

	// using lOOP

	firstName, lastName, email, userTicketsCount := getUserInput()
	// validations
	isValidName, isValidEmail, isValidTicket := helper.ValidateUserInputs(firstName, lastName, email, userTicketsCount, remainingTickets)

	// conditions
	if isValidName && isValidEmail && isValidTicket {
		bookTicket(userTicketsCount, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTicketsCount, firstName, lastName, email)

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
			// break
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
	wg.Wait()
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

		firstNames = append(firstNames, aBooking.firstName)

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
func bookTicket(userTicketsCount uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTicketsCount
	bookings[0] = firstName + " " + lastName

	//creating map data structure
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTicketsCount,
	}

	bookingSlice = append(bookingSlice, userData)

	fmt.Printf("List of booking is %v\n", bookingSlice)
	fmt.Printf("Thank you %v %v for booking %v tickets, You will receive a confirmation email at %v \n", firstName, lastName, userTicketsCount, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)                                                    // Put app to wait for 10 seconds before it sends the ticket
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) //Sprintf: format strings then you console them after
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v to email address %v \n", ticket, email)
	fmt.Println("################")
	wg.Done()
}
