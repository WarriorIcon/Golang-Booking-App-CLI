package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50

	fmt.Printf("Welcome to the %v booking application!\n", conferenceName) 
	fmt.Printf("We have %v tickets and %v are still available!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend!")

	var userName string
	var userTickets int = 2

	fmt.Printf("User %v booked %v tickets", userName, userTickets)
}



