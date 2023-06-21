package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("go routine 1: done")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("go routine 2: done")
	}()

	go func() {
		defer wg.Done()

		file, _ := os.CreateTemp("", "test.txt")
		_, _ = file.Write([]byte(strings.Repeat("a", 1_000_000_000)))

		fmt.Println("done writing")
	}()

	wg.Wait()
	time.Sleep(2 * time.Second)
}
