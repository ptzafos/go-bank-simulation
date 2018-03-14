package main

import (
	"time"
	"math/rand"
)

func SimulateClientToEntrance(e *Entrance) {

	for i := 0; i < B; i++ {
		e.newClient()
	}

}

func CustomersAreComing(e *[]Entrance) {

	for {
		for i := 0; i < A; i++ {
			// In case we want to pass the changes to main flow
			// That's why we use the pointers.
			go SimulateClientToEntrance(&((*e)[i]))
		}
		time.Sleep(GetRandomSleepTime(timeFrameForNewCustomers))
	}
}

func SimulateServeClient(s *ServePoint) {

	s.serveClient()
}

func ServePointStartWorking(s *[]ServePoint) {
	for {
		select {
		case <-syncQueue:
			for i := 0; i < C; i++ {
				// In case we want to pass the changes to main flow
				// That's why we use the pointers.
				go SimulateServeClient(&(*s)[i])
			}
		}
	}
}

func GetRandomSleepTime(tempo int) time.Duration {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	sleepTime := random.Int()%tempo + 1
	return time.Duration(sleepTime) * time.Millisecond
}
