package main

import (
	"fmt"
	"time"
)

func main() {
	asyncNotWorking()
	//asyncTmpFix()
}

func asyncNotWorking() {
	now := time.Now()

	go generateTask("task1", 100)
	go generateTask("task2", 200)
	go generateTask("task3", 0)
	go generateTask("task4", 100)

	fmt.Println("elapsed", time.Now().Sub(now))
}

func asyncTmpFix() {
	now := time.Now()

	go generateTask("task1", 100)
	go generateTask("task2", 200)
	go generateTask("task3", 0)
	go generateTask("task4", 100)

	fmt.Println("elapsed", time.Now().Sub(now))
	time.Sleep(time.Second)
}

func generateTask(text string, t time.Duration) {
	time.Sleep(t * time.Millisecond)
	fmt.Println(text)
}
