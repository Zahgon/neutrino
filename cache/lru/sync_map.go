package lru

import "sync"

// syncMap wraps a sync.Map with type parameters such that it's easier to
// access the items stored in the map since no type assertion is needed. It
// also requires explicit type definition when declaring and initiating the
// variables, which helps us understanding what's stored in a given map.
//
// NOTE: this is unexported to avoid confusion with `lnd`'s `SyncMap`.
type syncMap[K comparable, V any] struct {
	sync.Map
}

// Store puts an item in the map.
func (m *syncMap[K, V]) Store(key K, value V) { _ = "STUB: not implemented"; return }

// Load queries an item from the map using the specified key. If the item
// cannot be found, an empty value and false will be returned. If the stored
// item fails the type assertion, a nil value and false will be returned.
func (m *syncMap[K, V]) Load(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

// nolint: gocritic

// Delete removes an item from the map specified by the key.
func (m *syncMap[K, V]) Delete(key K) {
	_ = "STUB: not implemented"

	// LoadAndDelete queries an item and deletes it from the map using the
	// specified key.
	return
}

func (m *syncMap[K, V]) LoadAndDelete(key K) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// nolint: gocritic

// Range iterates the map.
func (m *syncMap[K, V]) Range(visitor func(K, V) bool) { _ = "STUB: not implemented"; return }

// Len returns the number of items in the map.
func (m *syncMap[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }
