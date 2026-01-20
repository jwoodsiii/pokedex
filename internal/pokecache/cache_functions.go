package pokecache

import (
	"errors"
	"time"
)

// Create a cache.Add() method that adds a new entry to the cache. It should take a key (a string) and a val (a []byte).
func (ch *Cache) Add(key string, val []byte) error {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	if key == "" || val == nil {
		return errors.New("need to supply a value for both key and val")
	}
	ch.entries[key] = cacheEntry{createdAt: time.Now().UTC(), val: val}
	return nil
}

// Create a cache.Get() method that gets an entry from the cache.
// It should take a key (a string) and return a []byte and a bool.
// The bool should be true if the entry was found and false if it wasn't.
func (ch *Cache) Get(key string) ([]byte, bool) {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	if val, ok := ch.entries[key]; ok == false {
		return nil, ok
	} else {
		return val.val, ok
	}
}

// Create a cache.reapLoop() method that is called when the cache is created
// (by the NewCache function). Each time an interval (the time.Duration passed to NewCache)
// passes it should remove any entries that are older than the interval.
// This makes sure that the cache doesn't grow too large over time.
// For example, if the interval is 5 seconds, and an entry was added 7 seconds ago,
// that entry should be removed.
func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for range ticker.C {
        c.reap(time.Now().UTC(), interval)
    }
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()
    cutoff := now.Add(-interval)
    for k, v := range c.entries {
        if v.createdAt.Before(cutoff) {
            delete(c.entries, k)
        }
    }
}
