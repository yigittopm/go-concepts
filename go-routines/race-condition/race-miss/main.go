package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	count := 0
	incrementFlag := flag.Bool("inc", false, "increment counter")
	decrementFlag := flag.Bool("dec", false, "decrement counter")
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if *incrementFlag {
			count++
		}
	}()

	go func() {
		defer wg.Done()
		if *decrementFlag {
			count--
		}
	}()
	wg.Wait()

	fmt.Println(count)
}
