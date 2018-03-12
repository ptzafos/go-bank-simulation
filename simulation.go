package main

import (
	"time"
	"fmt"
)

func SimulateClientToEntrance(e *Entrance){

	time.Sleep(time.Duration(GetRandomSleepTime(10000)) * time.Millisecond)
	fmt.Println(e.newClient(), e.entranceNum)

}


