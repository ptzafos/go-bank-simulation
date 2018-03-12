package main

import (
	"math/rand"
	"time"
)

//number of entrances
const A int = 100
const B int = 5
const C int = 5

func main() {

	start()
	e := make([]Entrance, A)
	for i := 0; i<A; i++{
		e[i].entranceNum = i+1
	}
	go CustomersAreComing(e)
	time.Sleep(1000000*time.Millisecond)
}

func CustomersAreComing(e []Entrance){

	for {
		threshold := GetRandomSleepTime(1000)
		for i := 0; i < A; i++{
			go SimulateClientToEntrance(&e[i])
		}
		time.Sleep(time.Duration(threshold) * time.Millisecond)
	}
}

func GetRandomSleepTime(tempo int) int{
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	sleepTime := random.Int()%tempo+1
	return sleepTime
}

