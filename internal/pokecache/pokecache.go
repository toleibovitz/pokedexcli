package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu sync.Mutex
	cacheMap map[string]cacheEntry
	interval time.Duration
}


type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	cacheMap := make(map[string]cacheEntry)
	cache := Cache{
		mu: sync.Mutex{},
		cacheMap: cacheMap,
		interval: interval,
	}
	go cache.reapLoop()
	return &cache

}
// Create a cache.Add() method that adds a new entry to the cache. It should take a key (a string) and a val (a []byte).
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.cacheMap[key] = entry
}
// Create a cache.Get() method that gets an entry from the cache. It should take a key (a string) and return a []byte and a bool. The bool should be true if the entry was found and false if it wasn't.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cacheMap[key]
	if !ok {
		return entry.val, false
	}
	return entry.val, true
}


// Create a cache.reapLoop() method that is called when the cache is created (by the NewCache function). Each time an interval (the time.Duration passed to NewCache) passes it should remove any entries that are older than the interval. This makes sure that the cache doesn't grow too large over time. For example, if the interval is 5 seconds, and an entry was added 7 seconds ago, that entry should be removed.
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		now := time.Now()
		c.mu.Lock()
		for key, value := range c.cacheMap {
			if now.Sub(value.createdAt) > c.interval {
				delete(c.cacheMap, key)
			}
		}
		c.mu.Unlock()
	}
}