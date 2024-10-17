package main
import "fmt"
func main(){

	var conferenceName = "Go conference"
	var conferenceTickets = 50
	var remainingTickets = 50;
	
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

}