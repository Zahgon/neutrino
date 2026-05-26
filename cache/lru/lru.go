package lru

import (
	"sync"

	"github.com/lightninglabs/neutrino/cache"
	"github.com/lightningnetwork/lnd/fn/v2"
)

// OnDeleteCallback is a function type that gets called when an element is
// deleted from the cache. It receives the key and value of the deleted element.
type OnDeleteCallback[K comparable, V cache.Value] func(key K, value V)

// entry represents a (key,value) pair entry in the Cache. The Cache's list
// stores entries which let us get the cache key when an entry is evicted.
type entry[K comparable, V cache.Value] struct {
	key   K
	value V
}

// Cache provides a generic thread-safe lru cache that can be used for
// storing filters, blocks, etc.
type Cache[K comparable, V cache.Value] struct {
	// capacity represents how much this cache can hold. It could be number
	// of elements or a number of bytes, decided by the cache.Value's Size.
	capacity uint64

	// size represents the size of all the elements currently in the cache.
	size uint64

	// ll is a doubly linked list which keeps track of recency of used
	// elements by moving them to the front.
	ll *List[entry[K, V]]

	// cache is a generic cache which allows us to find an elements position
	// in the ll list from a given key.
	cache syncMap[K, *Element[entry[K, V]]]

	// mtx is used to make sure the Cache is thread-safe.
	mtx sync.RWMutex

	// onDelete is a callback that is called when an element is deleted from
	// the cache.
	onDelete fn.Option[OnDeleteCallback[K, V]]
}

// CacheOption is a function that can be used to configure the cache.
type CacheOption[K comparable, V cache.Value] func(*Cache[K, V])

// WithDeleteCallback adds a delete callback to the cache which is called
// when an element is deleted from the cache.
func WithDeleteCallback[K comparable, V cache.Value](
	callback OnDeleteCallback[K, V]) CacheOption[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// NewCache return a cache with specified capacity, the cache's size can't
// exceed that given capacity.
func NewCache[K comparable, V cache.Value](capacity uint64,
	opts ...CacheOption[K, V]) *Cache[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Apply all options.

// evict will evict as many elements as necessary to make enough space for a new
// element with size needed to be inserted.
func (c *Cache[K, V]) evict(needed uint64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// We still need to evict some more elements.

// We should never reach here.

// Find the least recently used item.

// Determine lru item's size.

// Account for that element's removal in evicted and
// cache size.

// Call the onDelete callback if set for the element.

// Remove the element from the cache.

// Put inserts a given (key,value) pair into the cache. If the key already
// exists, it will replace value and update it to be most recent item in cache.
// The return value indicates whether items had to be evicted to make room for
// the new element.
func (c *Cache[K, V]) Put(key K, value V) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Load the element.

// Update the internal list inside a lock.

// If the element already exists, remove it and decrease cache's size.

// Then we need to make sure we have enough space for the element, evict
// elements if we need more space.

// We have made enough space in the cache, so just insert it.

// Release the lock.

// Update the cache.

// Get will return value for a given key, making the element the most recently
// accessed item in the process. Will return nil if the key isn't found.
func (c *Cache[K, V]) Get(key K) (V, error) { _ = "STUB: not implemented"; return *new(V), nil }

// Element not found in the cache.

// When the cache needs to evict a element to make space for another
// one, it starts eviction from the back, so by moving this element to
// the front, it's eviction is delayed because it's recently accessed.

// Len returns number of elements in the cache.
func (c *Cache[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

// Delete removes an item from the cache.
func (c *Cache[K, V]) Delete(key K) { _ = "STUB: not implemented"; return }

// LoadAndDelete queries an item and deletes it from the cache using the
// specified key.
func (c *Cache[K, V]) LoadAndDelete(key K) (V, bool) {
	_ = "STUB: not implemented"

	// Noop if the element doesn't exist.
	return *new(V), false
}

// Get its size.

// Call the onDelete callback if set for the element.

// Remove the element from the list and update the cache's size.

// Range iterates the cache without any ordering.
func (c *Cache[K, V]) Range(visitor func(K, V) bool) {
	_ = "STUB: not implemented"
	// valueVisitor is a closure to help unwrap the value from the cache.
	return
}

// RangeFILO iterates the items with FILO order, behaving like a stack.
func (c *Cache[K, V]) RangeFILO(visitor func(K, V) bool) { _ = "STUB: not implemented"; return }

// Stops the iteration if the visitor returns false to mimick
// the same behavior of `Range`.

// RangeFIFO iterates the items with FIFO order, behaving like a queue.
func (c *Cache[K, V]) RangeFIFO(visitor func(K, V) bool) { _ = "STUB: not implemented"; return }

// Stops the iteration if the visitor returns false to mimick
// the same behavior of `Range`.

// Size returns the total size of all elements in the cache. It uses
// the same units produced by V.Size().
func (c *Cache[K, V]) Size() uint64 { _ = "STUB: not implemented"; return 0 }
