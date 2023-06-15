package main

func main() {
	var receiveChan <-chan int
	var sendChan chan<- int
	intChan := make(chan int)

	receiveChan = intChan
	sendChan = intChan

	receiveChan = receiveChan
	sendChan = sendChan
}
