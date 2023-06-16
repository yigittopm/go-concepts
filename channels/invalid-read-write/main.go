package main

func main() {
	writeStream := make(chan<- int)
	readStream := make(<-chan int)

	//<-writeStream
	//readStream<- 1

	writeStream = writeStream
	readStream = readStream
}
