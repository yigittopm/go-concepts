package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	go fmt.Println("I am running in different thread")

	_, _ = http.Get("https://www.google.com")
}
