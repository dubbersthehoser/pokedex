package pokecache

import (
	"sync"
	"time"
)

const reapLoopTime time.Duration = 5 * time.Millisecond

type cacheEntry struct {
	val []byte
	createdAt time.Time
}

type Cache struct {
	entrys     map[string]cacheEntry
	mux      sync.Mutex
	duration time.Duration
}

func NewCache(t time.Duration) *Cache {
	c := &Cache{
		entrys: make(map[string]cacheEntry),
		mux: sync.Mutex{},
		duration: t,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, data []byte) {
	entry := cacheEntry{
		val: data,
		createdAt: time.Now(),
	}
	c.mux.Lock()
	c.entrys[key] = entry
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	entry, ok := c.entrys[key]
	c.mux.Unlock()
	if !ok {
		return nil, ok
	}
	return entry.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(reapLoopTime)
	defer ticker.Stop()
	for {
		<-ticker.C
		c.mux.Lock()
		for key, entry := range c.entrys {
			if c.duration <= time.Since(entry.createdAt) {
				delete(c.entrys, key)
			}
		}
		c.mux.Unlock()
	}
}
