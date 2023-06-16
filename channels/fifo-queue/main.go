package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan string, 2)
	people := []string{"John", "Jane", "Mike", "Steve", "Alex"}

	go func() {
		defer close(queue)

		for _, p := range people {
			time.Sleep(time.Second)
			fmt.Println(p, " starts waiting")
			queue <- p
		}
	}()

	for person := range queue {
		time.Sleep(2 * time.Second)
		fmt.Println("servering: ", person)
	}

	time.Sleep(time.Second)
	fmt.Println("closing the restaurant")
}
