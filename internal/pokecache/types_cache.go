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
	entries		map[string]cacheEntry
	mu			*sync.Mutex
	interval	time.Duration

}

// You'll probably want to expose a NewCache() function that creates a new cache with a configurable interval (time.Duration).
func NewCache(interval time.Duration) Cache {
	ch := Cache{
		entries: 	map[string]cacheEntry{},
		interval: 	interval,
	}
	ch.reapLoop()
	return ch
}
