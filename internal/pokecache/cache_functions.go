package pokecache

import (
	"fmt"
	"sync"
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
	ch.entries[key] = val
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
		return val, ok
	}
}

// Create a cache.reapLoop() method that is called when the cache is created
// (by the NewCache function). Each time an interval (the time.Duration passed to NewCache)
// passes it should remove any entries that are older than the interval.
// This makes sure that the cache doesn't grow too large over time.
// For example, if the interval is 5 seconds, and an entry was added 7 seconds ago,
// that entry should be removed.
func (ch *Cache) reapLoop() {
	ticker := time.NewTicker(ch.interval)
	defer ticker.Stop()
	go func() {
		ch.mu.Lock()
		defer ch.mu.Unlock()
		for t := range ticker.C {
			for k, v ch.entries {
				if time.Now().Sub(v.createdAt) >= ch.interval {
					delete(ch.entries, k)
				}
			}
		}
	}
}
