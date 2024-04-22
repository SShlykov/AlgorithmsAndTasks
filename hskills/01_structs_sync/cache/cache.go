package cache

import "sync"

type item[T any] struct {
	value T
}

type InMemoryCache[Key comparable, Val any] struct {
	mu    sync.RWMutex
	cache map[Key]item[Val]
}

func NewInMemoryCache[Key comparable, Val any]() *InMemoryCache[Key, Val] {
	return &InMemoryCache[Key, Val]{
		cache: make(map[Key]item[Val]),
	}
}

func (c *InMemoryCache[Key, Val]) Put(key Key, value Val) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = item[Val]{value: value}
}

func (c *InMemoryCache[Key, Val]) Get(key Key) (Val, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.cache[key]
	if ok {
		return v.value, true
	}

	return v.value, false
}

func (c *InMemoryCache[Key, Val]) GetOrCreate(key Key, value Val) Val {
	if v, ok := c.Get(key); ok {
		return v
	}

	c.Put(key, value)
	return value
}
