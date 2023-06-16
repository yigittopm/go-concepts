package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	timer := time.NewTimer(2 * time.Second)

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("tick: ", t)
		case <-timer.C:
			fmt.Println("time is up")
			return
		}
	}
}
