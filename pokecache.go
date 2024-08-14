package main

import (
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
func (c Cache) NewCache(interval time.Time) {
	for _, c := range c.cacheEntry {
		c.createdAt = interval.UTC()
	}
	c.reapLoop()
}
func (c Cache) Add(key string, val []byte) {

}
func (c Cache) Get(key string, val []byte, found bool) bool {
	return false
}
func (c Cache) reapLoop() {

}