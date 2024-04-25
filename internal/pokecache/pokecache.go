package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	return Cache{
		cache: make(map[string]cacheEntry),
	}
}

// adds into the cache
func (c *Cache) Add(key string, value []byte) {
	// add the values in the specific key
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

// gets the data from the cache
func (c *Cache) Get(key string) ([]byte, bool) {
	// reads the value from the passed in
	val, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return val.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

// cleans the data
func (c *Cache) reap(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(timeAgo) {
			delete(c.cache, k)
		}
	}
}
