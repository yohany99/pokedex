package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	mu       sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{cacheMap: make(map[string]cacheEntry)}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ce1 := cacheEntry{time.Now(), val}
	c.cacheMap[key] = ce1
}

func (c *Cache) Get(key string) ([]byte, bool) {
	val, ok := c.cacheMap[key]
	if ok {
		return val.val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {

	}
}
