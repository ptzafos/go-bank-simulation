package main

func main() {

	start()
	var block = make(chan int)
	bank := Bank{make([]ServePoint,C), make([]Entrance,A)}
	for i := 0; i < A; i++{
		bank.entrances[i].entranceNum = i+1
	}
	for i := 0; i < C; i++{
		bank.servePoints[i].servePointNum = i+1
	}
	go ServePointStartWorking(&bank.servePoints)
	go CustomersAreComing(&bank.entrances)
	for {
		select{
			case <-block:
		}
	}
}

