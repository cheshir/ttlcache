package main

import (
	"fmt"
	"time"

	"github.com/cheshir/ttlcache"
)

const (
	minute = 60 * time.Second
	hour   = 60 * minute
)

func Example() {
	// How often we need to stop the world and remove outdated records.
	resolution := minute

	cache := ttlcache.New(resolution)
	cache.Set(ttlcache.StringKey("some key"), "value", hour)
	value, ok := cache.Get(ttlcache.StringKey("some key"))
	if !ok {
		// there is no record with key "some key" in the cache. Probably it has been expired.
	}

	fmt.Println(value.(string)) // This is necessary type assertion, because returned value is of interface{} type.
	// Output: value
}
