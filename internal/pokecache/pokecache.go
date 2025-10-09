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

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.cacheMap[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cacheMap[key]
	if !ok {
		return entry.val, false
	}
	return entry.val, true
}



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