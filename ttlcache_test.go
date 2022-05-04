package ttlcache

import (
	"testing"
	"time"
)

func TestCache_GetSet(t *testing.T) {
	ttl := 2 * time.Millisecond
	key := StringKey("key")
	value := "value"

	c := New(time.Millisecond)
	c.Set(key, value, ttl)

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

	time.Sleep(4 * time.Millisecond)

	_, ok3 := c.Get(key)
	if ok3 {
		t.Error("record was not cleaned up")
	}
}

func TestCache_Delete(t *testing.T) {
	key := StringKey("key")
	value := "value"

	c := New(time.Second) // Cleanup should not be triggered.
	c.Set(key, value, 0)

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

	c.Delete(key)

	_, ok3 := c.Get(key)
	if ok3 {
		t.Error("record was not removed")
	}
}

func TestCache_Clear(t *testing.T) {
	c := New(time.Millisecond)

	for i := 1; i < 5; i++ {
		c.Set(IntKey(i), i, 0)
	}

	c.Clear()

	for i := 1; i < 5; i++ {
		_, ok := c.Get(IntKey(i))
		if ok {
			t.Error("Storage was not cleaned up")
		}
	}

	// Verify that the cleanup manager is still running
	ttl := 2 * time.Millisecond
	key := StringKey("key")
	value := "value"

	c.Set(key, value, ttl)
	time.Sleep(4 * time.Millisecond)

	_, ok := c.Get(key)
	if ok {
		t.Error("record was not cleaned up")
	}
}

func TestClose(t *testing.T) {
	c := New(time.Second)
	c.Set(IntKey(1), 1, 0)
	c.Set(IntKey(2), 2, 0)
	c.Set(IntKey(3), 3, 0)
	c.Set(IntKey(4), 4, 0)

	err := c.Close()
	if err != nil {
		t.Error("Unexpected error")
	}

	_, ok := c.Get(IntKey(1))
	if ok {
		t.Error("Storage was not cleaned up")
	}
}
