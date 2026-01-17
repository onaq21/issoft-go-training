package cache

import (
	"sync"
	"time"
	"errors"
)

type Cache struct {
	capacity int
	items map[string]CacheElem
	mu sync.Mutex
}

type CacheElem struct {
	value any
	lastAccess time.Time
}

func CreateCache(cap int) (*Cache, error) {
	if cap < 1 { return nil, errors.New("Cache create error: capacity must be a positive number") }
	return &Cache{cap, map[string]CacheElem{}, sync.Mutex{}}, nil
}

func (cache *Cache) Set(key string, obj any) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	
	if len(cache.items) == cache.capacity {
	oldestKey := ""
	oldestTime := time.Now()

	for k, value := range cache.items {
		if value.lastAccess.Before(oldestTime) {
			oldestKey = k
			oldestTime = value.lastAccess
		}
	}
		delete(cache.items, oldestKey)
	}

	cache.items[key] = CacheElem{obj, time.Now()}
}

func (cache *Cache) Get(key string) any {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	el, ok := cache.items[key]
	if !ok { return nil }

	el.lastAccess = time.Now()
	cache.items[key] = el
	return el.value
}

func (cache *Cache) Remove(key string) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	_, ok := cache.items[key]
	if !ok { return false }
	delete(cache.items, key)
	return true
}

func (cache *Cache) Scan() {
	for {
		cache.mu.Lock()
		t := time.Now()
		for key, val := range cache.items {
			if t.Sub(val.lastAccess) > time.Second * 10 {
				delete(cache.items, key)
			}
		}
		cache.mu.Unlock()
		time.Sleep(time.Second * 1)
	}
}