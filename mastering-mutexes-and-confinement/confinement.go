package main

import (
	"fmt"
	"sync"
)

func manageTicket(ticketChannel chan int, doneChannel chan bool, exitChannel chan bool, tickets *int) {
	for {
		select {
		case user := <-ticketChannel:
			if *tickets > 0 {
				*tickets--
				fmt.Printf("User %d purchased a ticket. Tickets remaining: %d\n", user, *tickets)
			} else {
				fmt.Printf("User %d found no tickets\n", user)
			}
		case <-doneChannel:
			fmt.Printf("Remaining tickets are: %d", *tickets)
			exitChannel <- true
			return
		}
	}
}

func buyTicketConfinement(wg *sync.WaitGroup, ticketChannel chan int, userId int) {
	defer wg.Done()
	ticketChannel <- userId
}

func main() {
	var wg sync.WaitGroup
	tickets := 500

	ticketChannel := make(chan int)
	doneChannel := make(chan bool)
	exitChannel := make(chan bool)

	go manageTicket(ticketChannel, doneChannel, exitChannel, &tickets)
	for userId := 0; userId < 200; userId++ {
		wg.Add(1)
		go buyTicketConfinement(&wg, ticketChannel, userId)
	}
	wg.Wait()

	doneChannel <- true
	<-exitChannel
}
