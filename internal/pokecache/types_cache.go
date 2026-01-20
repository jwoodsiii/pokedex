package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

type Cache struct {
    entries 	map[string]cacheEntry
    mu   		sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
    c := &Cache{
        entries: make(map[string]cacheEntry),
    }
    go c.reapLoop(interval)
    return c
}
