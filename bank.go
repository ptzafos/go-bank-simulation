package main

import(
	"fmt"
	"github.com/rs/xid"
	"sync"
	"sync/atomic"
)

var entranceLock sync.Mutex

var queue_num uint32

func start(){
	fmt.Println("Simulation started, welcome to UOM bank.")
}

type Bank struct {

	serve_points []Serve_point
	entrances []Entrance

}



type Entrance struct {

	clients []Client

}

func (e *Entrance) newClient() Client{

	client := Client{genXid(),take_number(&queue_num)}
	entranceLock.Lock()
	e.clients = append(e.clients, client)
	entranceLock.Unlock()
	return client
}

type Client struct{

	client_id string
	ticket_num uint32

}

func genXid() string{
	id := xid.New()
	return id.String()
}

type Serve_point struct{

}

type Waiting_queue struct{


}

func take_number(queue_num *uint32) uint32{
	atomic.AddUint32(queue_num, 1)
	return *queue_num;
}

