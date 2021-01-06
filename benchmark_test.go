package ttlcache

import (
	"testing"
	"time"
)

// The Set time is highly depends on the map capacity.
// The main problem is map expansion.
func BenchmarkCache_Set(b *testing.B) {
	resolution := 9999 * time.Second // Avoid stopping the world.
	uniqueKeysCount := 100
	c := New(resolution)

	for i := 0; i < b.N; i++ {
		c.Set(IntKey(i%uniqueKeysCount), i, 0)
	}
}

func BenchmarkCache_Get(b *testing.B) {
	resolution := 9999 * time.Second // Avoid stopping the world.
	uniqueKeysCount := 100
	c := New(resolution)

	for i := 0; i < 100; i++ {
		c.Set(IntKey(i), i, 0)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		val, _ := c.Get(IntKey(i % uniqueKeysCount))
		_ = val
	}
}
