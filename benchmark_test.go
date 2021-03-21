package ttlcache

import (
	"testing"
	"time"
)

// The Set time is highly depends on the map capacity.
// The main problem is map expansion.
func BenchmarkCache_Set_100(b *testing.B) {
	benchmarkCacheSet(b, 100)
}

func BenchmarkCache_Set_1000(b *testing.B) {
	benchmarkCacheSet(b, 1000)
}

func BenchmarkCache_Set_10000(b *testing.B) {
	benchmarkCacheSet(b, 10000)
}

func BenchmarkCache_Get_100(b *testing.B) {
	benchmarkCacheGet(b, 100)
}

func BenchmarkCache_Get_1000(b *testing.B) {
	benchmarkCacheGet(b, 1000)
}

func BenchmarkCache_Get_10000(b *testing.B) {
	benchmarkCacheGet(b, 10000)
}

func benchmarkCacheSet(b *testing.B, numberOfKeys int) {
	resolution := 9999 * time.Second // Avoid stopping the world.
	c := New(resolution)

	for i := 0; i < b.N; i++ {
		c.Set(IntKey(i%numberOfKeys), i, 0)
	}
}

func benchmarkCacheGet(b *testing.B, numberOfKeys int) {
	resolution := 9999 * time.Second // Avoid stopping the world.
	c := New(resolution)

	for i := 0; i < 100; i++ {
		c.Set(IntKey(i), i, 0)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		val, _ := c.Get(IntKey(i % numberOfKeys))
		_ = val
	}
}
