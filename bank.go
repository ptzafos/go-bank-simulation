package main

import(
	"fmt"
	"sync"
)

var entranceLock sync.Mutex
var queueNum uint32
//var customerQueue [100]chan Client
var customerQueue = make(chan Client, 10000)

func start(){
	fmt.Println("Simulation started, welcome to UOM bank.")
}

type Bank struct {

	serve_points []ServePoint
	entrances []Entrance

}

type Entrance struct {

	entranceNum int

}

func (s *ServePoint) serveClient(){


}

func (e *Entrance) newClient() {

	entranceLock.Lock()
	client := Client{take_number(&queueNum)}
	fmt.Println("Customer arrived", client.ticketNum, e.entranceNum)
	customerQueue <- client
	entranceLock.Unlock()

}

type Client struct{
	ticketNum uint32
}

type ServePoint struct{


}


func take_number(queue_num *uint32) uint32{
	//No need to use as we already lock the add operation in the MUTEX
	//atomic.AddUint32(queue_num, 1)
	queueNum+=1
	return *queue_num
}

