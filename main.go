package main
import "fmt"
func main(){

	var conferenceName = "Go conference" // conferenceName :="Go conference" // means the same thing and it doesn't work on the constants and variable that are explicitly 
	var conferenceTickets = 50
	var remainingTickets uint = 50;
	
	fmt.Println("Welcome to", conferenceName,"booking application")
	fmt.Println("We have a total of", conferenceTickets, "ticket and ", remainingTickets ,"are still available.")
	fmt.Println("Get your tickets here to attend")

	// print formatting
	fmt.Printf("Welcome to %v booking application \n remaining ticket %v \n ", conferenceName, remainingTickets)

	var userName string
	var userTickets int

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

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address : ")
	fmt.Scan(&email)


	fmt.Println("Enter number of tickets : ")
	fmt.Scan(&userTicketsCount)
	remainingTickets = remainingTickets - userTicketsCount

	fmt.Printf("Thank you %v %v for booking %v tickets, You will receive a confirmation email at %v \n", firstName, lastName, userTicketsCount, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}