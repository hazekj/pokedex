package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	pokeMap map[string]cacheEntry
	mu      *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	pokeCache := Cache{
		pokeMap: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
	}
	go pokeCache.reapLoop(interval)
	return pokeCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.pokeMap[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.pokeMap[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		currentTime := time.Now()
		for key, entry := range c.pokeMap {
			diff := currentTime.Sub(entry.createdAt)
			if diff > interval {
				c.mu.Lock()
				delete(c.pokeMap, key)
				fmt.Println("Deleted cache", key)
				c.mu.Unlock()
			}
		}
	}
}
