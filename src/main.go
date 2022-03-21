package main

import (
	"fmt"
	"strings"
)

func main() {
	goConf := "Go!-Conf"
	const capacity uint8 = 50
	var ticketsAvailable uint8 = 35

	fmt.Printf("Welcome to %s\n", goConf)
	fmt.Printf("Total visitor capacity: %d\n", capacity)
	fmt.Printf("Tickets available: %d\n", ticketsAvailable)

	// var bookings []string // Slice type
	bookings := []string{} // Slice type

	// var bookings [capacity]string // Array type

	for {

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

		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)
		fmt.Printf("Thank You!! %v %v for booking %v tickets. You will receive confirmation over email at %v\n", firstName, lastName, userTickets, email)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("Our bookings: %v\n", bookings)
		fmt.Printf("First name of all bookings: %v\n", firstNames)
	}

	// fmt.Printf("%v tickets remaining for %v!\n", ticketsAvailable, goConf)
}
