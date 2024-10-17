package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go conference" // conferenceName :="Go conference" // means the same thing and it doesn't work on the constants and variable that are explicitly
	var conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings [50]string
	//var bookingSlice []string //
	bookingSlice := []string{}

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "ticket and ", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	// print formatting
	fmt.Printf("Welcome to %v booking application \n remaining ticket %v \n ", conferenceName, remainingTickets)

	var userName string
	var userTickets uint

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets. \n", userName, userTickets)
	fmt.Printf("datatype is %T and value is %v \n", userName, userName)
	fmt.Scan(&userName)
	fmt.Printf("%v isn stored memory location =%v  \n", userName, &userName)

	// address
	var firstName string
	var lastName string
	var email string
	var userTicketsCount uint //accept only positive integers

	// using lOOP
	for {

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last Name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address : ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets : ")
		fmt.Scan(&userTicketsCount)
		remainingTickets = remainingTickets - userTicketsCount
		if userTickets < remainingTickets {
			bookings[0] = firstName + " " + lastName
			bookingSlice = append(bookingSlice, firstName+" "+lastName)

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

			fmt.Printf("Thank you %v %v for booking %v tickets, You will receive a confirmation email at %v \n", firstName, lastName, userTicketsCount, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, aBooking := range bookingSlice {
				var names = strings.Fields(aBooking)
				firstNames = append(firstNames, names[0])

			}
			fmt.Printf("The first names of booking are: %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Our conference is sold out please come next year")
				break
			}

		} else if userTickets == remainingTickets {

			fmt.Printf(" you are buying every ticket %v tickets remaining, so you can't book %v tickets \n", remainingTickets, userTickets)

		}else {

			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets \n", remainingTickets, userTickets)
			continue //continue to the next iteration
		}
	}

}
