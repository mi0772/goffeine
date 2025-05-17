# Goffeine

**Goffeine** is a high-performance in-memory caching library for Go, inspired by Java's [Caffeine](https://github.com/ben-manes/caffeine).

Designed to be lightweight and efficient, Goffeine offers fast access, automatic eviction, and expiration mechanisms with clean and simple APIs.

---

## Features

- Fast in-memory caching
- Least Recently Used (LRU) eviction policy
- Time-To-Live (TTL) expiration per entry
- Optional default TTL for all entries
- Configurable maximum number of entries
- Future support for TinyLFU, count-min sketch and advanced eviction
- Clean API with no external dependencies

---

## Installation

```bash
go get github.com/mi0772/goffeine
```

---

## Example usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/mi0772/goffeine"
)

func main() {
	cache := goffeine.NewCache(goffeine.Config{
		MaxEntries: 100,
		DefaultTTL: 5 * time.Second,
	})

	cache.Put("user:42", "Walter White")
	val, ok := cache.Get("user:42")
	if ok {
		fmt.Println("Found:", val)
	}

	time.Sleep(6 * time.Second)
	_, ok = cache.Get("user:42")
	fmt.Println("Still exists after TTL?", ok) // false
}
```

---

## API Overview

```go
type Config struct {
MaxEntries int           // maximum number of elements in the cache
DefaultTTL time.Duration // default TTL for entries (optional)
}

func NewCache(cfg Config) *Cache

func (c *Cache) Put(key string, value any)
func (c *Cache) Get(key string) (any, bool)
func (c *Cache) Remove(key string)
func (c *Cache) Len() int
```

---

## Roadmap

| Feature                      | Status         |
|------------------------------|----------------|
| Basic LRU cache              | In Progress    |
| Entry TTL support            | Planned        |
| Sharded concurrency support  | Planned        |
| TinyLFU & Count-Min Sketch   | Planned        |
| Usage statistics             | Planned        |
| Prometheus metrics export    | Planned        |

---

## Why Goffeine?

- You want a fast in-memory cache but don't want to run Redis.
- You need something smarter than `sync.Map` and simpler than BigCache.
- You want to use a Caffeine-like system, but in native Go.

---

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

---

## Author

Carlo Di Giuseppe  
GitHub: [https://github.com/mi0772](https://github.com/mi0772)

---

## Contributions

Contributions, feature requests, and issues are welcome.  
Goffeine is in active development.
