package ttlcache

import (
	"sync"
	"time"
)

const defaultCapacity = 16 // Just to avoid extra allocations in most of the cases.

// Cache represents key-value storage.
type Cache struct {
	done  chan struct{}
	mu    sync.RWMutex
	items map[uint64]item
}

type item struct {
	deadline int64 // Unix nano
	value    interface{}
}

// New creates key-value storage.
// resolution – configures cleanup manager.
// Cleanup operation locks storage so think twice before setting it to small value.
func New(resolution time.Duration) *Cache {
	c := &Cache{
		done:  make(chan struct{}),
		items: make(map[uint64]item, defaultCapacity),
	}

	go cleaner(c, resolution)

	return c
}

// Get returns stored record.
// The first returned variable is a stored value.
// The second one is an existence flag like in the map.
func (c *Cache) Get(key uint64) (interface{}, bool) {
	c.mu.RLock()
	cacheItem, ok := c.items[key]
	c.mu.RUnlock()

	if !ok {
		return nil, false
	}

	return cacheItem.value, true
}

// Set adds value to the cache with given ttl.
// ttl value should be a multiple of the resolution time value.
func (c *Cache) Set(key uint64, value interface{}, ttl time.Duration) {
	cacheItem := item{
		deadline: time.Now().UnixNano() + int64(ttl),
		value:    value,
	}

	c.mu.Lock()
	c.items[key] = cacheItem
	c.mu.Unlock()
}

// Delete removes record from storage.
func (c *Cache) Delete(key uint64) {
	c.mu.Lock()
	delete(c.items, key)
	c.mu.Unlock()
}

// Clear removes all items from storage and leaves the cleanup manager running.
func (c *Cache) Clear() {
	c.mu.Lock()
	c.items = make(map[uint64]item, defaultCapacity)
	c.mu.Unlock()
}

// Close stops cleanup manager and removes records from storage.
func (c *Cache) Close() error {
	close(c.done)

	c.mu.Lock()
	c.items = nil
	c.mu.Unlock()

	return nil
}

// cleanup removes outdated items from the storage.
// It triggers stop the world for the cache.
func (c *Cache) cleanup() {
	now := time.Now().UnixNano()
	c.mu.Lock()

	for key, item := range c.items {
		if item.deadline < now {
			delete(c.items, key)
		}
	}

	c.mu.Unlock()
}

func cleaner(c *Cache, resolution time.Duration) {
	ticker := time.NewTicker(resolution)

	for {
		select {
		case <-ticker.C:
			c.cleanup()
		case <-c.done:
			ticker.Stop()
			return
		}
	}
}
