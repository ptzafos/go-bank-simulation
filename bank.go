package main

import(
	"fmt"
	"sync"
	"time"
)

//number of entrances
const A = 2
//number of customers coming in random time frames
const B = 3
//number of serve points
const C = 5
//Max Bank Capacity
const D = 100
//Adding zeros after one changes the speed of the simulation(makes it slower)
const timeWaitDegree = 10
//Time from a servePoint to serve a customer degree in Millis
const customerServeTime = 1000 * timeWaitDegree
//Time Frame for new Customers to come to an entrance in in Millis
const timeFrameForNewCustomers = 1300 * timeWaitDegree


var entranceLock sync.Mutex
var queueNum int
var customerQueue = make(chan Client, D)
//Help channel in order to use for select in ServePointsSimulation
var syncQueue = make(chan int, D)

func start(){
	fmt.Println("Simulation started, welcome to UOM bank.")
	fmt.Println()
	fmt.Println("Clients\t\t\t\t\tQueueNum-WaitTime\t\t\t\tServePoint-NumServing\t\tServePoint-NumServed")
	fmt.Println("------------------------------------------------------------------------------------------------------")
}

type Bank struct {

	servePoints []ServePoint
	entrances []Entrance

}

type Entrance struct {

	entranceNum int

}

func (e *Entrance) newClient() {

	entranceLock.Lock()
	client := Client{takeNumber(&queueNum)}
	syncQueue <- 1
	customerQueue <- client
	fmt.Println("New Customer Arrived\tTicket:", queueNum, "AvWaitTime" , timeWaitDegree*float32(len(customerQueue))/float32(C), "sec")
	entranceLock.Unlock()
}

func (s *ServePoint) serveClient(){

	s.ServePointLock.Lock()
	client := <- customerQueue
	fmt.Println("\t\t\t\t\t\t\t\t\t\t\t\t\t\tServPoint:",s.servePointNum, "TickServ:",client.ticketNum)
	time.Sleep(GetRandomSleepTime(customerServeTime))
	fmt.Println("\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\tServPoint:",s.servePointNum, "TickServed:",client.ticketNum)
	s.ServePointLock.Unlock()
}


type Client struct{
	ticketNum int
}

type ServePoint struct{

	ServePointLock sync.Mutex
	servePointNum int

}

func takeNumber(queueNum *int) int{
	//No need to use as we already lock the add operation in the MUTEX
	//atomic.AddUint32(queue_num, 1)
	*queueNum+=1
	return *queueNum
}

