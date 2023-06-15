package critical

import (
	"strconv"
	"testing"
)

func BenchmarkCriticalMutexCache(b *testing.B) {
	cache := newMutexCache()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		go cache.set("key", "value"+strconv.Itoa(i))
	}
}

func BenchmarkCriticalChannelCache(b *testing.B) {
	cache := newChannelCache()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		go cache.set("key", "value"+strconv.Itoa(i))
	}
}