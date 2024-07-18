package cache

// Cache is a simple key / value cache
type Cache interface {
	Get(key string) (any, bool)
	Set(key string, value any)
	Del(key string)
	Size() int
	Clear() int
}

type memoryCache struct {
	cache map[string]any
}

// NewMemoryCache creates a new Cache
func NewMemoryCache() Cache {
	return &memoryCache{
		cache: make(map[string]any),
	}
}

// Get will try to get associated with a key from the cache, if present
func (c *memoryCache) Get(key string) (any, bool) {
	v, ok := c.cache[key]
	return v, ok
}

// Set will store a key-value pair in the cache
func (c *memoryCache) Set(key string, value any) {
	c.cache[key] = value
}

// Del will remove a key and its associated value from the cache.
func (c *memoryCache) Del(key string) {
	delete(c.cache, key)
}

// Size will return the current number of entries in the cache.
func (c *memoryCache) Size() int {
	return len(c.cache)
}

// Clear will empty the cache, removing all stored key-value pairs.
func (c *memoryCache) Clear() int {
	count := len(c.cache)
	c.cache = make(map[string]any)
	return count
}
