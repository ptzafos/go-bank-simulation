package main

import (
	"time"
)


func SimulateClientToEntrance(e *Entrance){

	time.Sleep(time.Duration(GetRandomSleepTime(100)) * time.Millisecond)
	for i:=0; i < B; i++{
		e.newClient()
	}
}


