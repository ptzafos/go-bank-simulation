package main

import (
	"os"
	"os/signal"
)

func main() {

	start()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	bank := Bank{make([]ServePoint,C), make([]Entrance,A)}
	for i := 0; i < A; i++{
		bank.entrances[i].entranceNum = i+1
	}
	for i := 0; i < C; i++{
		bank.servePoints[i].servePointNum = i+1
	}
	go ServePointStartWorking(&bank.servePoints)
	go CustomersAreComing(&bank.entrances)
	<-interrupt
}

