package main

import "sync"

func StartSimulation(e *Entrance, wg *sync.WaitGroup){
	defer wg.Done()
	e.newClient()
}


