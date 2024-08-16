package main

import (
	"reflect"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}
type Cache struct {
	cacheEntry map[string]cacheEntry
	mu *sync.Mutex
}
func NewCache(interval time.Duration) {
	c := Cache{
		cacheEntry: make(map[string]cacheEntry),
		mu:   &sync.Mutex{},
	}
	go c.reapLoop(interval)
}
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cE := cacheEntry{createdAt: time.Now(), val: val}
	c.cacheEntry[key] = cE
}
func (c *Cache) Get(key string, val []byte, found bool) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return reflect.DeepEqual(c.cacheEntry[key].val, val)
}
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
}
}
func (c *Cache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k := range c.cacheEntry {
		now := time.Since(c.cacheEntry[k].createdAt)
		if  now > interval{
			delete(c.cacheEntry,k)
		}
	}
}