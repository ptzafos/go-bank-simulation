package main

import (
	"math/rand"
	"time"
	"fmt"
)

//number of entrances
const A int = 2
//number of customers coming in random time frames
const B int = 3
//number of serve points
const C int = 5

func main() {

	start()
	e := make([]Entrance, A)
	//var e *[A]Entrance
	for i := 0; i<A; i++{
		e[i].entranceNum = i+1
	}
	go CustomersAreComing(&e)
	time.Sleep(10000*time.Millisecond)
	for {
		select{
		case c := <- customerQueue:
			fmt.Println("Client:", c)
		default:
			return
		}

	}
}

func CustomersAreComing(e *[]Entrance){

	for {
		threshold := 1000
		for i := 0; i < A; i++{
			// In case we want to pass the changes to main flow if we had a more complex Entrance Struct.
			// That's why we use the pointers.
			go SimulateClientToEntrance(&((*e)[i]))
		}
		time.Sleep(time.Duration(threshold) * time.Millisecond)
	}
}

func ServePointStartWorking(s []ServePoint){

}

func GetRandomSleepTime(tempo int) int{
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	sleepTime := random.Int()%tempo+1
	return sleepTime
}

