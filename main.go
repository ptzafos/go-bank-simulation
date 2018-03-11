package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	start()
	var e Entrance
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go StartSimulation(&e, &wg)
	}
	wg.Wait()
	fmt.Println(len(e.clients))
	//time.Sleep(3000 * time.Millisecond)

}

