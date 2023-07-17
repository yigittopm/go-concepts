package main

import (
	"fmt"
	"sync"
)

type Button struct {
	Clicked *sync.Cond
}

func main() {
	button := Button{
		Clicked: sync.NewCond(&sync.Mutex{}),
	}

	subscribe := func(c *sync.Cond, param string, fn func(s string)) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)

		go func(p string) {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()

			c.Wait()

			fn(p)
		}(param)

		goroutineRunning.Wait()
	}

	var clickedRegistered sync.WaitGroup

	for _, v := range []string{
		"Maximizin window",
		"Displaying annoying dialog box!",
		"Mouse clicked.",
	} {
		clickedRegistered.Add(1)

		subscribe(button.Clicked, v, func(s string) {
			fmt.Println(s)
			clickedRegistered.Done()
		})
	}

	button.Clicked.Broadcast()

	clickedRegistered.Wait()

}
