// Package goffeine provides an in-memory caching system inspired by Caffeine (Java),
// with support for eviction policies like LRU and TTL expiration.
package goffeine

// Cache defines the basic behavior of an in-memory cache
// with asynchronous and synchronous put operations.
type Cache interface {
	// PutAndForget inserts a value associated with the given key,
	// without waiting for internal processing (e.g. eviction, TTL cleanup).
	PutAndForget(key string, value any)

	// PutAndWait inserts a value and waits until the operation is fully applied.
	PutAndWait(key string, value any)

	// Get retrieves the value associated with the given key.
	// Returns (nil, false) if the key is not found or expired.
	Get(key string) (any, bool)

	// Remove deletes the key from the cache. Returns true if the key was present.
	Remove(key string) bool

	// Len returns the number of currently stored entries.
	Len() int
}

// CacheOf defines a type-safe variant of the Cache interface,
// using generics to enforce value type safety.
type CacheOf[V any] interface {
	// PutAndForget inserts a value associated with the given key,
	// without waiting for internal processing.
	PutAndForget(key string, value V)

	// PutAndWait inserts a value and waits until the operation is applied.
	PutAndWait(key string, value V)

	// Get retrieves the value of type V for the given key.
	Get(key string) (V, bool)

	// Remove deletes the key from the cache. Returns true if it existed.
	Remove(key string) bool

	// Len returns the total number of entries in the cache.
	Len() int
}
