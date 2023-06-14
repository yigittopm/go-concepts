package benchmarks

import (
	"fmt"
	"sync"
	"testing"
)

var (
	bufferedWrites   int
	unbufferedWrites int
)

func TestMain(m *testing.M) {
	m.Run()
	fmt.Println("buffered writes: ", bufferedWrites)
	fmt.Println("un-buffered writes: ", unbufferedWrites)
}

func BenchmarkUnbufferedChannel(b *testing.B) {
	b.ReportAllocs()

	done := make(chan any)
	defer close(done)

	ch := make(chan int)
	go reader(done, ch)

	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			unbufferedWrites++
			ch <- 1
		}(i)
	}
	wg.Wait()
}

func BenchmarkBufferedChannel(b *testing.B) {
	b.ReportAllocs()
	done := make(chan any)
	defer close(done)

	bufferSize := 5000
	ch := make(chan int, bufferSize)
	go reader(done, ch)

	var size int
	if b.N-bufferSize < 0 {
		size = b.N
	} else {
		size = bufferSize
	}

	var wg sync.WaitGroup
	for i := 0; i < b.N; i += size {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < size; j++ {
				bufferedWrites++
				ch <- i + j
			}
		}(i)
	}
	wg.Wait()
}

func reader(done chan any, in <-chan int) {
	for {
		select {
		case <-done:
			return
		case <-in:
		}
	}
}
