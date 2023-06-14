package benchmarks

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"testing"
)

func BenchmarkUnbufferedWrite(b *testing.B) {
	b.ReportAllocs()
	write(b, newTmpFile())
}

func BenchmarkBufferedWrite(b *testing.B) {
	b.ReportAllocs()
	write(b, bufio.NewWriter(newTmpFile()))
}

func newTmpFile() *os.File {
	file, err := ioutil.TempFile("", "tmp")
	if err != nil {
		log.Fatalf("could not creat temporary file: %v", err)
	}

	return file
}

func write(b *testing.B, w io.Writer) {
	done := make(chan any)
	defer close(done)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := w.Write([]byte(strconv.Itoa(i)))
		if err != nil {
			log.Fatalf("could not write to file: %v", err)
		}
	}
}
