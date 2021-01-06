package ttlcache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	ttl := 2 * time.Millisecond
	key := StringKey("key")
	value := "value"

	c := New(time.Millisecond)
	c.Set(key, value, ttl)

	{
		val, ok := c.Get(key)
		if !ok {
			t.Error("storage missed expected value")
		}

		v, ok2 := val.(string)
		if !ok2 {
			t.Error("type assertion failed")
		}

		if v != value {
			t.Errorf("incorret value: got: %v expected: %v", v, value)
		}
	}

	time.Sleep(4 * time.Millisecond)

	{
		_, ok := c.Get(key)
		if ok {
			t.Error("value was not cleaned up from storage")
		}
	}
}
