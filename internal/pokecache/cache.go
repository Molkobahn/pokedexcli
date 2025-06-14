package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	entry		map[string]cacheEntry
	mu			sync.Mutex
	interval	time.Duration
}

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

func NewCache(inter time.Duration) *Cache {
	cache := Cache {
		entry:		make(map[string]cacheEntry),
		interval:	inter,
	}

	go cache.reapLoop()

	return &cache
}


func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	
	currentTime := time.Now()
	newEntry := cacheEntry {
		createdAt:	currentTime,
		val:		val,
	}

	c.entry[key] = newEntry

	c.mu.Unlock()
}


func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	
	defer c.mu.Unlock()

	entry, exists := c.entry[key]
	if exists {
		return entry.val, true
	} else {
		return nil, false
	}

}


func (c *Cache) reapLoop() {
	
	ticker := time.NewTicker(c.interval)

	for range ticker.C {
		
		cutoffTime := time.Now().Add(-c.interval)
		tempSlice := []string{}

		for key, value := range c.entry {
			if value.createdAt.Before(cutoffTime) {
				tempSlice = append(tempSlice, key)
				
			}
		}
		
		func() {
			c.mu.Lock()
			defer c.mu.Unlock()
			for _, slice := range tempSlice {
				delete(c.entry, slice)
			}
		}()
	}
}