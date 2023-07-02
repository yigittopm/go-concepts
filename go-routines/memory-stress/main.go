package main

func main() {
	done := make(chan struct{})

	go func() {
		memory := make([]string, 0)
		for {
			memory = append(memory, "hello")
		}
	}()

	<-done
}
