package main

import "fmt"

func main() {
	goConf := "Go!-Conf"
	const capacity uint8 = 50
	var ticketsAvailable uint8 = 35

	fmt.Printf("Welcome to %s\n", goConf)
	fmt.Printf("Total visitor capacity: %d\n", capacity)
	fmt.Printf("Tickets available: %d\n", ticketsAvailable)

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Print("Enter your first name?: ")
	fmt.Scan(&firstName) // Scanning requires pointers in GoLang

	fmt.Print("Enter your last name?: ")
	fmt.Scan(&lastName) // Scanning requires pointers in GoLang

	fmt.Print("Enter your email?: ")
	fmt.Scan(&email) // Scanning requires pointers in GoLang

	fmt.Print("How many tickets do you want?: ")
	fmt.Scan(&userTickets)

	ticketsAvailable -= uint8(userTickets)

	fmt.Printf("Thank You!! %v %v for booking %v tickets. You will receive confirmation over email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v!\n", ticketsAvailable, goConf)

}
