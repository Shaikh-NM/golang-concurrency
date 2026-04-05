package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func buyTicketMutex(wg *sync.WaitGroup, userId int, remainingTickets *int) {
	defer wg.Done()

	mutex.Lock()
	defer mutex.Unlock() // it is best practice to unlock using defer as this ensures even if the function returns early/crashes then the mutex is released. this prevents deadlock

	if *remainingTickets > 0 {
		*remainingTickets--
		fmt.Printf("User %d purchased a ticket. Tickets remaining %d\n", userId, *remainingTickets)
	} else {
		fmt.Printf("User %d found no ticket.\n", userId)
	}
}

func main() {
	var tickets int = 500
	var wg sync.WaitGroup

	for userId := 0; userId < 500; userId++ {
		wg.Add(1)
		go buyTicketMutex(&wg, userId, &tickets)
	}

	wg.Wait()
}
